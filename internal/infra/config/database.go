package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDBClient(ctx context.Context) *pgxpool.Pool {
	if pool, err := pgxpool.NewWithConfig(ctx, Config()); err != nil {
		_ = fmt.Errorf(err.Error())
	} else {
		return pool
	}
	return &pgxpool.Pool{}
}
