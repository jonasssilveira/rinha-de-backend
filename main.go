package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	dbInstance, err := sql.Open("postgres", BuildURLConnection())
	if err != nil {
		log.Fatalf("Fail to open connection with database: %v", err)
	}

	dbInstance.SetConnMaxLifetime(0)
	dbInstance.SetMaxIdleConns(1)
	dbInstance.SetMaxOpenConns(1)

	db = dbInstance
}

func main() {
	defer db.Close()

	store := NewPostgresTransactionStore(db)
	server := NewServer(store)

	addr := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	log.Printf("Listening in %s...", addr)

	err := server.Handler.Listen(":8080")
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Fail to start server on addr: %q", addr)
	}
}
