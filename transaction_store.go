package main

import (
	"context"
)

type TransactionStore interface {
	AddClient(clientId int, balance, limit int) error
	GetBalance(clientId int) (ClientBalance, error)
	UpdateBalance(clientId int, clientBalance ClientBalance) error
	AddTransactionAsync(
		ctx context.Context,
		clientId int,
		transaction Transaction,
		processTransaction func(c ClientBalance, t Transaction) (ClientBalance, error),
	) (ClientBalance, error)
	GetTransactions(clientId, count int) ([]Transaction, error)
}
