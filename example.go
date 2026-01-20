package main

import "fmt"

func main() {
	name := "Semen"
	//a := 1
	//b := 2
	//c := a + b
	//fmt.Println("Sum:", c)
	fmt.Printf("age: %s\n", name)
}

//Задание №1
// go fmt .\example.go

//Задание №2
//go vet ./...
//# Learning_EM
//# [Learning_EM]
//.\example.go:11:2: fmt.Printf format %d has arg name of wrong type string
//после исправления
//go vet ./...
//пусто...
