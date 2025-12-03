package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // драйвер PostgreSQL
)

func main() {
	dsn := "host=localhost port=5432 user=postgres password=password dbname=app_db sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка открытия базы:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Ошибка подключения к базе:", err)
	}

	fmt.Println("Подключение успешно!")
}
