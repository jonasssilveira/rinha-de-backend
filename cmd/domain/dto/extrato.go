package dto

import (
	db "rinha-de-backend-2024-q1/db/sqlc"
	"time"
)

type Extrato struct {
	OutSaldo   `json:"saldo"`
	Transacoes []db.GetClienteTrasacoesRow `json:"ultimas_transacoes"`
}

type OutSaldo struct {
	Total       int64     `json:"total"`
	DataExtrato time.Time `json:"data_extrato"`
	Limite      int64     `json:"limite"`
}
