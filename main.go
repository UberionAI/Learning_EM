package main

import "fmt"

func SumElementSlice(s []int) int {
	sum := 0
	for _, a := range s {
		sum += a
	}
	return sum
}

func main() {
	s := []int{1, 2, 3}
	fmt.Println(SumElementSlice(s))
}
