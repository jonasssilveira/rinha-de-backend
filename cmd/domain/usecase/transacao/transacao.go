package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

	defer func(err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}(err)

	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return dto.Saldo{}, fiber.NewError(http.StatusNotFound, "not found")
		}

		fmt.Println(err.Error())
		return dto.Saldo{}, fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	switch transacaoDTO.Tipo {
	case "d":
		if (saldo.Valor - transacaoDTO.Valor) < 0 {
			return dto.Saldo{}, fiber.NewError(http.StatusUnprocessableEntity)
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
	default:
		return dto.Saldo{}, fiber.NewError(http.StatusUnprocessableEntity)
	}

	if err != nil {
		fmt.Println(err.Error())
		return dto.Saldo{}, err
	}

	err = t.clientRepository.CreateTransacoes(ctx, db.CreateTransacoesParams{
		ClienteID: id,
		Valor:     transacaoDTO.Valor,
		Tipo:      transacaoDTO.Tipo,
		Descricao: transacaoDTO.Descricao,
	})

	if err != nil {
		fmt.Println(err.Error())
		return dto.Saldo{}, fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	return dto.Saldo{
		Limite: saldo.Limite.Int64,
		Saldo:  saldo.Valor - transacaoDTO.Valor,
	}, nil
}
