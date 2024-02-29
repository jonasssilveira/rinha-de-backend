package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

type DatabaseConn struct {
	SGDatabase   string
	Login        string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

func BuildURLConnection() string {
	databaseURLConnection := DatabaseConn{
		SGDatabase:   "postgresql",
		Login:        os.Getenv("LOGIN_DATABASE"),
		Password:     os.Getenv("PASSWORD_DATABASE"),
		Host:         os.Getenv("HOST_DATABASE"),
		Port:         os.Getenv("PORT_DATABASE"),
		DatabaseName: "rinha",
	}

	return fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=disable",
		databaseURLConnection.SGDatabase, databaseURLConnection.Login,
		databaseURLConnection.Password, databaseURLConnection.Host,
		databaseURLConnection.Port, databaseURLConnection.DatabaseName)

}

func Config() *pgxpool.Config {
	databaseUrl := BuildURLConnection()
	fmt.Println(databaseUrl)
	dbConfig, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		exec := c.QueryRow(ctx, "SELECT count(*) FROM pg_stat_activity;")
		var quantConn int
		exec.Scan(&quantConn)
		fmt.Printf("Quantidade de conexoes: %d\n", quantConn)

		return true
	}
	return dbConfig
}
