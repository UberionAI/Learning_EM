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

	//	Задание №2
	//rows, err := db.Query("SELECT id, username FROM users")
	//if err != nil {
	//	log.Fatal("query error: ", err)
	//}
	//defer rows.Close()
	//
	//for rows.Next() {
	//	var id int
	//	var username string
	//	if err := rows.Scan(&id, &username); err != nil {
	//		log.Fatal("error reading:", err)
	//	}
	//
	//	fmt.Printf("id:%d name:%s\n", id, username)
	//}

	//Задание №3
	//username := "Alex"
	//rows, err := db.Query("SELECT id FROM users WHERE username = $1", username)
	//if err != nil {
	//	log.Fatal("query error: ", err)
	//}
	//defer rows.Close()
	//fmt.Printf("user with username: %s\n", username)
	//for rows.Next() {
	//	var id int
	//
	//	if err := rows.Scan(&id); err != nil {
	//		log.Fatal("error reading:", err)
	//	}
	//	fmt.Printf("id:%d name:%s\n", id, username)
	//}

	// Задание №4
	//	newUser := "Gregory"
	//	var newId int
	//	err = db.QueryRow(`
	//        INSERT INTO users (username) VALUES ($1)
	//        RETURNING id`,
	//		newUser).Scan(&newId)
	//
	//	if err != nil {
	//		log.Fatal("Error inserting:", err)
	//	}
	//	fmt.Printf("Добавлен: %s (ID: %d)\n", newUser, newId)
	//

	// Задание №5:
	newAge := 25
	result, err := db.Exec(`
        UPDATE users 
        SET age = $1 
        WHERE username = $2`,
		newAge, "Alex")

	if err != nil {
		log.Fatal("Error update:", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated rows: %d\nAlex is now %d \n", rowsAffected, newAge)
}
