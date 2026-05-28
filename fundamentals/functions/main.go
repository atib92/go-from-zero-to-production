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

// function that returns multiple values: the swapped strings
func swap(a string, b string) (string, string) {
	return b, a
}

// variadic function that sums an arbitrary number of integers
func variadic_sum(numbers ...int) int {
	total := 0

	for _, num := range numbers {
		total += num
	}

	return total
}

// closure example:  A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables; in this sense, the function is "bound" to the variables.
func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
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

	fmt.Println("Variadic Sum:", variadic_sum(1, 2, 3, 4, 5))

	count := counter()
	fmt.Println("Counter:", count()) // 1
	fmt.Println("Counter:", count()) // 2
	fmt.Println("Counter:", count()) // 3
}
