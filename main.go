package main

import (
	"fmt"
	"os"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	_ = Add(1, 2)

	if err := os.Remove("non-existing.txt"); err != nil {
		fmt.Println("remove failed", err)
	}
}

//Задание №10
//Первый запуск: errcheck ./...
//main.go:15:14:  fmt.Fprintln(os.Stdout, "hello")
//main.go:16:11:  os.Remove("non-existent-file") // <-- errcheck

//После исправления:
//errcheck ./...
//Пусто....
