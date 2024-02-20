package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	usecaseclient "rinha-de-backend-2024-q1/cmd/domain/usecase/cliente"
	usecasetransacao "rinha-de-backend-2024-q1/cmd/domain/usecase/transacao"
	db "rinha-de-backend-2024-q1/db/sqlc"
	"rinha-de-backend-2024-q1/internal/infra/config"
	"rinha-de-backend-2024-q1/internal/infra/handlers"
)

func main() {
	app := fiber.New()

	database := config.GetDBClient(context.Background())

	dbClient := db.New(database)
	defer database.Reset()

	useCaseClientSaldo := usecaseclient.NewClientInfo(dbClient)
	useCaseClientTransacao := usecasetransacao.NewTransacao(dbClient)

	handleTransacao := handlers.NewTransacao(useCaseClientTransacao)
	handleSaldo := handlers.NewSaldoHandler(useCaseClientSaldo)

	app.Post("/clientes/:id/transacoes", handleTransacao.HandleCreateTransacao)
	app.Get("/clientes/:id/extrato", handleSaldo.GetExtrato)

	err := app.Listen(":8080")
	if err != nil {
		fmt.Printf("error to start server, error %v", err.Error())
	}
}
