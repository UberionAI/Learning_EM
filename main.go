package main

import "fmt"

func ChangeValue(ptr *int) {
	*ptr = 100
}

func main() {
	a := 10
	fmt.Println("До:", a)
	ChangeValue(&a)
	fmt.Println("Новое значение:", a)
}
