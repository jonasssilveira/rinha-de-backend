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
)

type TransacaoClient interface {
	CreateClientTrasacao(ctx context.Context, id int32, transacaoDTO dto.Transacao) (dto.Saldo, error)
}

type Transacao struct {
	clientRepository repository.Client
}

func NewTransacao(
	clientRepository repository.Client,
) *Transacao {
	return &Transacao{
		clientRepository: clientRepository,
	}
}

func (t Transacao) CreateClientTrasacao(ctx context.Context, id int32, transacaoDTO dto.Transacao) (dto.Saldo, error) {
	var err error
	saldo, err := t.clientRepository.GetSaldoCliente(ctx, id)

	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return dto.Saldo{}, fiber.NewError(http.StatusNotFound, "not found")
		}
		return dto.Saldo{}, fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	switch transacaoDTO.Tipo {
	case "d":

		if (saldo.Limite.Int64 - (transacaoDTO.Valor - saldo.Valor)) < 0 {
			return dto.Saldo{}, fiber.NewError(http.StatusUnprocessableEntity, "valor da transacao nao pode diminuir o limite abaixo de 0")
		}

		err = t.clientRepository.Withdraw(ctx, db.WithdrawParams{
			Valor:     transacaoDTO.Valor,
			ClienteID: id,
		})
	case "c":
		err = t.clientRepository.Deposit(ctx, db.DepositParams{
			Valor:     transacaoDTO.Valor,
			ClienteID: id,
		})
	}

	if err != nil {
		return dto.Saldo{}, err
	}

	err = t.clientRepository.CreateTransacoes(ctx, db.CreateTransacoesParams{
		ClienteID: id,
		Valor:     transacaoDTO.Valor,
		Tipo:      transacaoDTO.Tipo,
		Descricao: transacaoDTO.Descricao,
	})

	if err != nil {
		return dto.Saldo{}, fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	return dto.Saldo{
		Limite: saldo.Limite.Int64,
		Saldo:  saldo.Valor - transacaoDTO.Valor,
	}, nil
}
