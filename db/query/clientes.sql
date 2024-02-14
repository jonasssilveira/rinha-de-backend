-- name: CreateCliente :one
INSERT INTO clientes (nome,
                 limite)
VALUES ($1, $2)RETURNING nome, limite;

-- name: GetCliente :one
SELECT limite, nome
FROM clientes
WHERE id = $1 LIMIT 1;

-- name: DeleteCliente :exec
DELETE
FROM clientes
WHERE id = $1;

