package main

import (
	"Learning_EM/task_1/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/users/", handlers.UsersHandler)
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil) //nolint:errcheck
}
