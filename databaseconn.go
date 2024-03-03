package main

import (
	"fmt"
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
