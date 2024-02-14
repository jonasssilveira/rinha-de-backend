package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

func GetDBClient() *sql.DB {
	DB, err := sql.Open(dbDriver, BuildURLConnection())
	if err != nil {
		_ = fmt.Errorf(err.Error())
		return nil
	}
	return DB
}
