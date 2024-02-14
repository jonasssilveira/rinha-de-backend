package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	usecase "rinha-de-backend-2024-q1/cmd/domain/usecase/cliente"
	"strconv"
)

type SaldoHandler struct {
	extratoClient usecase.ClientExtrato
}

func NewSaldoHandler(client usecase.ClientExtrato) *SaldoHandler {
	return &SaldoHandler{extratoClient: client}
}

func (s *SaldoHandler) GetExtrato(ctx *fiber.Ctx) error {
	idClient, err := strconv.ParseInt(ctx.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	extrato, err := s.extratoClient.GetClientExtrato(ctx.Context(), int32(idClient))
	if err != nil {
		return err
	}

	return ctx.JSON(extrato)
}
