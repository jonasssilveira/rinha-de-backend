-- Coloque scripts iniciais aqui

--Criação das tabelas
CREATE TABLE clientes
(
    id     SERIAL PRIMARY KEY,
    nome   VARCHAR(50) NOT NULL,
    limite BIGSERIAL   NOT NULL
);

CREATE TABLE transacoes
(
    id           SERIAL PRIMARY KEY,
    cliente_id   INTEGER     NOT NULL,
    valor        BIGSERIAL   NOT NULL,
    tipo         CHAR(1)     NOT NULL,
    descricao    VARCHAR(255) NOT NULL,
    realizada_em TIMESTAMP   NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_clientes_transacoes_id
        FOREIGN KEY (cliente_id) REFERENCES clientes (id)
);

CREATE TABLE saldos
(
    id         SERIAL PRIMARY KEY,
    cliente_id INTEGER   NOT NULL,
    valor      BIGSERIAL NOT NULL,
    CONSTRAINT fk_clientes_saldos_id
        FOREIGN KEY (cliente_id) REFERENCES clientes (id)
);

-- Create index on cliente_id column
CREATE INDEX idx_cliente_id ON transacoes (cliente_id);

-- Create index on realizada_em column
CREATE INDEX idx_realizada_em ON transacoes (realizada_em);

--Criação da function
CREATE
OR REPLACE FUNCTION check_table_size()
RETURNS TRIGGER AS $$
DECLARE
table_size INT;
    max_table_size
INT := 5;
BEGIN
-- Execute the query to get the table size and store it in table_size variable
EXECUTE 'SELECT COUNT(*) FROM ' || TG_RELNAME::regclass INTO table_size;
IF
table_size >= max_table_size THEN
        RAISE EXCEPTION 'Table size limit reached';
END IF;
RETURN NEW;
END;
$$
LANGUAGE plpgsql;

--Inserção da function
CREATE TRIGGER enforce_table_size_limit
    BEFORE INSERT
    ON clientes
    FOR EACH ROW
    EXECUTE FUNCTION check_table_size();

INSERT INTO clientes (nome, limite)
VALUES ('o barato sai caro', 1000 * 100),
       ('zan corp ltda', 800 * 100),
       ('les cruders', 10000 * 100),
       ('padaria joia de cocaia', 100000 * 100),
       ('kid mais', 5000 * 100);
INSERT INTO transacoes (cliente_id, valor, tipo, descricao, realizada_em)
VALUES (1, 500, 'D', 'Deposit', '2023-12-31 10:30:00'),
       (1, -200, 'W', 'Withdrawal', '2023-12-30 09:45:00'),
       (2, 1000, 'D', 'Deposit', '2023-12-29 15:20:00'),
       (3, -300, 'W', 'Withdrawal', '2023-12-28 14:10:00'),
       (4, 800, 'D', 'Deposit', '2023-12-27 11:00:00');
INSERT INTO saldos (cliente_id, valor)
VALUES (1, 300),
       (2, 1500),
       (3, 800),
       (4, 1200),
       (5, 2000);