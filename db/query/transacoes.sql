-- name: CreateTransacoes :exec
INSERT INTO transacoes (cliente_id,
                      valor,
                      tipo,
                      descricao)
VALUES ($1, $2, $3, $4);

-- name: GetClienteTrasacoes :many
SELECT t.valor, t.tipo, t.descricao, t.realizada_em
FROM transacoes t
LEFT JOIN clientes c on c.id = t.cliente_id
WHERE c.id = $1 order by t.realizada_em desc LIMIT 10;
