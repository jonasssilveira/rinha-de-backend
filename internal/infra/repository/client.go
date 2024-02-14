package repository

import (
	"context"
	db "rinha-de-backend-2024-q1/db/sqlc"
)

type Client interface {
	CreateCliente(ctx context.Context, arg db.CreateClienteParams) (db.CreateClienteRow, error)
	DeleteCliente(ctx context.Context, id int32) error
	GetCliente(ctx context.Context, id int32) (db.GetClienteRow, error)
	Deposit(ctx context.Context, arg db.DepositParams) error
	Withdraw(ctx context.Context, arg db.WithdrawParams) error
	CreateSaldo(ctx context.Context, arg db.CreateSaldoParams) error
	GetSaldoCliente(ctx context.Context, id int32) (db.GetSaldoClienteRow, error)
	CreateTransacoes(ctx context.Context, arg db.CreateTransacoesParams) error
	GetClienteTrasacoes(ctx context.Context, id int32) ([]db.GetClienteTrasacoesRow, error)
}
