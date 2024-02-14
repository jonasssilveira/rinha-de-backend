package mocks

import (
	"context"
	db "rinha-de-backend-2024-q1/db/sqlc"
)

type DBClient struct {
	MockCreateCliente func(ctx context.Context, arg db.CreateClienteParams) (db.CreateClienteRow, error)
	MockDeleteCliente func(ctx context.Context, id int32) error
	MockGetCliente    func(ctx context.Context, id int32) (db.GetClienteRow, error)
}

func (db *DBClient) CreateCliente(ctx context.Context, arg db.CreateClienteParams) (db.CreateClienteRow, error) {
	return db.MockCreateCliente(ctx, arg)
}

func (db *DBClient) DeleteCliente(ctx context.Context, id int32) error {
	return db.MockDeleteCliente(ctx, id)
}

func (db *DBClient) GetCliente(ctx context.Context, id int32) (db.GetClienteRow, error) {
	return db.MockGetCliente(ctx, id)
}
