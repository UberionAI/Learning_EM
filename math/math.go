package math

import "fmt"

func Sum(a int, b int) int {
	return a + b
}

func SumOnlyNatural(a int, b int) int {
	if a < 0 || b < 0 {
		fmt.Errorf("%d and %d must be non-negative", a, b)
	}
	return (a + b)
}
