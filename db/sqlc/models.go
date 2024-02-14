// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"
)

type Cliente struct {
	ID     int32  `json:"id"`
	Nome   string `json:"nome"`
	Limite int64  `json:"limite"`
}

type Saldo struct {
	ID        int32 `json:"id"`
	ClienteID int32 `json:"cliente_id"`
	Valor     int64 `json:"valor"`
}

type Transaco struct {
	ID          int32     `json:"id"`
	ClienteID   int32     `json:"cliente_id"`
	Valor       int64     `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadaEm time.Time `json:"realizada_em"`
}