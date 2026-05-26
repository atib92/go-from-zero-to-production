package main

import (
	"errors"
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func swap(a string, b string) (string, string) {
	return b, a
}

func main() {
	sum := add(10, 20)
	fmt.Println("Sum:", sum)

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Division Result:", result)

	first, second := swap("Go", "Python")

	fmt.Println("First:", first)
	fmt.Println("Second:", second)
}