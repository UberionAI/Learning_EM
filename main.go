package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Age      uint
}

func main() {
	// Задание №7
	//Установим ГОРМ и драйвер для горма
	dsn := "host=localhost port=5432 user=postgres password=password dbname=app_db sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	newUser := User{Username: "VovaGORM", Age: 30}
	db.Create(&newUser)
	fmt.Printf("User %s created with ID %v and age %d", newUser.Username, newUser.ID, newUser.Age)

	var user User
	db.First(&user, "username = ?", "VovaGORM")
	fmt.Printf("\nUser %s found with ID %v", user.Username, user.ID)

	user.Age = 40
	db.Save(&user)
	fmt.Printf("\nUser %s updated to age %d", user.Username, user.Age)

	db.Delete(&user)
	fmt.Printf("\nUser %s deleted", user.Username)
}
