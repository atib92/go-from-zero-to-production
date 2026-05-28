package main

import "fmt"

func printArray(numbers [5]int) {
	fmt.Println("Inside Function:", numbers)
}

func modifyArray(numbers [5]int) {
	numbers[0] = 999
}

func main() {
	var numbers [5]int

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	fmt.Println("Initial Array:", numbers)

	fruits := [3]string{"apple", "banana", "orange"}

	fmt.Println("Fruits:", fruits)

	fmt.Println("Array Length:", len(fruits))

	printArray(numbers)

	modifyArray(numbers)

	fmt.Println("After modifyArray():", numbers)

	copied := numbers

	copied[0] = 500

	fmt.Println("Original Array:", numbers)
	fmt.Println("Copied Array:", copied)
}