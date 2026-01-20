package main

import "fmt"

func main() {
	fmt.Println("hello")
}

//Задание 7
//PS GolandProjects\Learning_EM> staticcheck ./...
//main.go:5:7: const unusedConst is unused (U1000)
//main.go:7:6: func unusedFunc is unused (U1000)

//Когда убрали лишнее и применили staticcheck ./... повторно
//пусто...
