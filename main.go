package main

import (
	"errors"
	"fmt"
)

func DivisionNum(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

func main() {
	res, err := DivisionNum(1.0, 2.0)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат:", res)
}
