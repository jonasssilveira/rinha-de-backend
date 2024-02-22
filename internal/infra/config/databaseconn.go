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

	return fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=disable&connect_timeout=%d",
		databaseURLConnection.SGDatabase, databaseURLConnection.Login,
		databaseURLConnection.Password, databaseURLConnection.Host,
		databaseURLConnection.Port, databaseURLConnection.DatabaseName, time.Duration(2))

}

func Config() *pgxpool.Config {
	databaseUrl := BuildURLConnection()
	fmt.Println(databaseUrl)
	dbConfig, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = int32(10000)
	dbConfig.MinConns = int32(20)
	dbConfig.MaxConnLifetime = time.Duration(500) * time.Second
	dbConfig.MaxConnIdleTime = time.Duration(500) * time.Millisecond
	dbConfig.ConnConfig.ConnectTimeout = time.Duration(10) * time.Second

	//dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
	//	log.Println("Before acquiring the connection pool to the database!!")
	//	return true
	//}
	//
	//dbConfig.AfterRelease = func(*pgx.Conn) bool {
	//	log.Println("After release the connection pool to the database!!")
	//	return true
	//}
	//
	//dbConfig.BeforeClose = func(c *pgx.Conn) {
	//	log.Println("Closed the connection pool to the database!!")
	//}

	return dbConfig
}
