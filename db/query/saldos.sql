-- name: CreateSaldo :exec
INSERT INTO saldos (cliente_id,
                      valor)
VALUES ($1, $2);

-- name: GetSaldoCliente :one
SELECT c.limite, s.valor
FROM saldos s
LEFT JOIN clientes c on c.id = s.cliente_id
WHERE c.id = $1 LIMIT 1;

-- name: Deposit :exec
UPDATE saldos
SET valor = valor + $1
WHERE cliente_id = $2;

-- name: Withdraw :exec
UPDATE saldos
SET valor = valor - $1
WHERE cliente_id = $2;