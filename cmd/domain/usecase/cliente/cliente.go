package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"rinha-de-backend-2024-q1/cmd/domain/dto"
	db "rinha-de-backend-2024-q1/db/sqlc"
	"rinha-de-backend-2024-q1/internal/infra/repository"
	"time"
)

type ClientExtrato interface {
	GetClientExtrato(ctx context.Context, id int32) (dto.Extrato, error)
	CreateClient(ctx context.Context, nome string, limit int64) (db.CreateClienteRow, error)
	DeleteClient(ctx context.Context, id int32) error
}

type ClientInfo struct {
	clientRepository repository.Client
}

func NewClientInfo(clientRepository repository.Client) *ClientInfo {
	return &ClientInfo{
		clientRepository: clientRepository,
	}
}

func (c ClientInfo) GetClientExtrato(ctx context.Context, id int32) (dto.Extrato, error) {
	transacoes, err := c.clientRepository.GetClienteTrasacoes(ctx, id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return dto.Extrato{}, fiber.NewError(http.StatusNotFound, "not found")
		}
		return dto.Extrato{}, fiber.NewError(http.StatusInternalServerError, err.Error())

	}
	saldo, err := c.clientRepository.GetSaldoCliente(ctx, id)
	if err != nil {
		return dto.Extrato{}, fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	return dto.Extrato{
		OutSaldo: dto.OutSaldo{
			Total:       saldo.Valor,
			DataExtrato: time.Now(),
			Limite:      saldo.Limite.Int64,
		},
		Transacoes: transacoes,
	}, nil
}

func (c ClientInfo) CreateClient(ctx context.Context, nome string, limit int64) (db.CreateClienteRow, error) {
	return c.clientRepository.CreateCliente(ctx, db.CreateClienteParams{
		Nome:   nome,
		Limite: limit,
	})
}

func (c ClientInfo) DeleteClient(ctx context.Context, id int32) error {
	return c.clientRepository.DeleteCliente(ctx, id)
}
