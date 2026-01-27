package main

import "fmt"

func main() {
	a := 50
	if a%2 == 0 {
		fmt.Printf("%d is even\n", a)
	} else {
		fmt.Printf("%d is odd\n", a)
	}
}
