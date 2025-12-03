package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	//Postgre поднимаем из docker-compose.yml
	//docker compose up -d

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

	rows, err := db.Query("SELECT id, username FROM users")
	if err != nil {
		log.Fatal("query error: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string
		if err := rows.Scan(&id, &username); err != nil {
			log.Fatal("error reading:", err)
		}

		fmt.Printf("id:%d name:%s\n", id, username)
	}
}
