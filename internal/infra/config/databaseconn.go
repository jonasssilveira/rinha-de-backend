package config

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"time"
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

	dbConfig.MaxConns = int32(1000)
	dbConfig.MinConns = int32(10)
	dbConfig.MaxConnLifetime = time.Duration(20) * time.Millisecond
	dbConfig.MaxConnIdleTime = time.Duration(20) * time.Millisecond

	return dbConfig
}
