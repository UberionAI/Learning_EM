package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is empty")
	}

	var db *sql.DB
	var err error
	for i := 0; i < 30; i++ {
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
			db.Close()
		}
		log.Printf("Wait Postgres... attempt %d/30: %v", i+1, err)
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		log.Fatalf("Cannot connect to Postgres after 30 sec: %v", err)
	}
	defer db.Close()

	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "dev"
	}
	fmt.Printf("Go App v%s started!\n", version)

	var pgVersion string
	err = db.QueryRow("SELECT version()").Scan(&pgVersion)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	fmt.Println("Connected to Postgres!")
	fmt.Println("Postgres version:", pgVersion)
}
