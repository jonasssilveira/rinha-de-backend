package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	MaxStatementTranscations        = 10
	MaxTransactionDescriptionLength = 10
)

var ErrDebitBelowLimit = errors.New("insufficient limit for this debit")
var ErrInvalidTransaction = errors.New("invalid transaction payload")

type Server struct {
	transactionStore TransactionStore
	Handler          *fiber.App
}

func NewServer(store TransactionStore) *Server {
	var server = new(Server)

	server.transactionStore = store
	server.Handler = setupRoutes(server)

	return server
}

func setupRoutes(server *Server) *fiber.App {
	app := fiber.New()
	app.Post("/clientes/:id/transacoes", server.transactions)
	app.Get("/clientes/:id/extrato", server.getExtrato)

	return app
}

func (s *Server) transactions(ctx *fiber.Ctx) error {
	clientId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity)
	}

	transaction, err := getTransactionFromBody(ctx.Body())
	if err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity)
	}
	transaction.TransactionDate = time.Now()

	clientBalance, err := s.addTransaction(ctx.Context(), clientId, transaction)
	if err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity)
	}

	return ctx.JSON(clientBalance)
}

func (s Server) addTransaction(ctx context.Context, clientId int, transaction Transaction) (ClientBalance, error) {
	return s.transactionStore.AddTransactionAsync(
		ctx,
		clientId,
		transaction,
		processTransaction,
	)
}

func (s *Server) getExtrato(ctx *fiber.Ctx) error {
	clientId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity)
	}
	if clientId > 5 {
		return fiber.NewError(http.StatusNotFound)
	}
	balanceChan := make(chan ClientBalance, 1)
	transactionsChan := make(chan []Transaction, 1)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		wg.Done()
		goRoutinebalance, err := s.transactionStore.GetBalance(clientId)
		if err != nil {
			return
		}
		balanceChan <- goRoutinebalance
	}()
	go func() {
		defer wg.Done()
		goRoutinetransactions, err := s.transactionStore.GetTransactions(clientId, MaxStatementTranscations)
		if err != nil {
			return
		}
		transactionsChan <- goRoutinetransactions
	}()
	wg.Wait()
	balance := <-balanceChan
	transactions := <-transactionsChan

	return ctx.JSON(buildStatement(balance, transactions))
}

func processTransaction(
	clientBalance ClientBalance,
	transaction Transaction,
) (ClientBalance, error) {
	if !isValidTransaction(transaction) {
		return clientBalance, ErrInvalidTransaction
	}

	switch transaction.Type {
	case TypeCredit:
		clientBalance.Balance += transaction.Amount
	case TypeDebit:
		newBalance := clientBalance.Balance - transaction.Amount
		if newBalance < -clientBalance.AccountLimit {
			return clientBalance, ErrDebitBelowLimit
		}

		clientBalance.Balance = newBalance
	}

	return clientBalance, nil
}

func isValidTransaction(t Transaction) bool {
	if t.Amount <= 0 {
		return false
	}

	if t.Description == "" {
		return false
	}

	if len(t.Description) > MaxTransactionDescriptionLength {
		return false
	}

	if t.Type != TypeCredit && t.Type != TypeDebit {
		return false
	}

	return true
}

func buildStatement(balance ClientBalance, transactions []Transaction) ClientStatement {
	return ClientStatement{
		Balance: ClientStatementBalance{
			Total:         balance.Balance,
			AccountLimit:  balance.AccountLimit,
			StatementDate: time.Now(),
		},
		LatestTransactions: transactions,
	}
}

func getTransactionFromBody(body []byte) (Transaction, error) {
	var transaction Transaction
	if err := json.Unmarshal(body, &transaction); err != nil {
		return transaction, err
	}
	return transaction, nil
}
