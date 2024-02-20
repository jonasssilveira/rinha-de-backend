package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"regexp"
	"rinha-de-backend-2024-q1/cmd/domain/dto"
	usecase "rinha-de-backend-2024-q1/cmd/domain/usecase/transacao"
	"strconv"
)

type Transacao struct {
	extratoClient usecase.TransacaoClient
}

func NewTransacao(clientTransacao usecase.TransacaoClient) *Transacao {
	return &Transacao{
		extratoClient: clientTransacao,
	}
}

func (t *Transacao) HandleCreateTransacao(ctx *fiber.Ctx) error {
	idClient, err := strconv.ParseInt(ctx.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	body := ctx.Body()
	var transacaoDTO dto.Transacao
	if err := json.Unmarshal(body, &transacaoDTO); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity)
	}

	if transacaoDTO.Descricao == "" || &transacaoDTO.Descricao == nil || regexp.MustCompile(`\s`).MatchString(transacaoDTO.Descricao) {
		return fiber.NewError(422)
	}

	if transacaoDTO.Valor == 0 {
		return fiber.NewError(422)
	}

	trasacao, err := t.extratoClient.CreateClientTrasacao(ctx.Context(), int32(idClient), transacaoDTO)

	if err != nil {
		return err
	}

	return ctx.JSON(trasacao)

}
