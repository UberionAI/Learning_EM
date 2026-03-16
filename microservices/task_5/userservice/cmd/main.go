package main

import (
	"fmt"
	"net/http"
	"userservice/internal/handlers"
)

func main() {
	http.HandleFunc("/users", handlers.UsersHandler)
	fmt.Println("UserService is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
