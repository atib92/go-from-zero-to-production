package main

import (
	"errors"
	"fmt"
)

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func findUser(id int) (string, error) {
	if id != 100 {
		return "", fmt.Errorf("user %d not found", id)
	}

	return "Atib", nil
}

func main() {
	result, err := divide(10, 2)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Result:", result)

	user, err := findUser(200)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User:", user)
}