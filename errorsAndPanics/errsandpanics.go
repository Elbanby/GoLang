package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Demonstrates errors and panics differences")
	fmt.Println("Errors: ")
	if _, err := precheckdivide(2, 0); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Panics: ")
	divide(2, 0)
}

func precheckdivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can't divide by zero!")
	}
	return divide(a, b), nil
}

func divide(a, b int) int {
	return a / b
}
