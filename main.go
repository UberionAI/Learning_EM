package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	p := Person{Name: "Vova", Age: 18, City: "New York"}
	fmt.Println(p)
}
