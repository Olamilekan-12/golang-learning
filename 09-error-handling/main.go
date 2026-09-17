package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("cannot divide by zero")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func processOrder(id int) error {
	_, err := divide(100, 0)
	if err != nil {
		return fmt.Errorf("Processing order %d: %w", id, err)
	}
	return nil
}

func main() {
	result, err := divide(10, 2)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Result: ", result)

	result2, err2 := divide(10, 0)

	if err2 != nil {
		fmt.Println("Error :", err2)

	}
	fmt.Println("Result :", result2)

	err3 := processOrder(42)
	if errors.Is(err3, ErrDivideByZero) {
		fmt.Println("the root cause was a divide by zero")
	}

}
