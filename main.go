package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: app <name>")
		os.Exit(1)
	}

	name := os.Args[1]
	fmt.Println("hello,", name)
}

//Задание 9
//module Learning_EM
//
//go 1.22
//
//require (
//	github.com/google/uuid v1.6.0
//)
//пишем go mod tidy для удаления ненужной зависимости из go.mod
//стало:
//module Learning_EM
//
//go 1.22
