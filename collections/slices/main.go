package main

import "fmt"

func modifySlice(numbers []int) {
	numbers[0] = 999
}

func appendValue(numbers []int) []int {
	numbers = append(numbers, 100)
	return numbers
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Original Slice:", numbers)

	modifySlice(numbers)

	fmt.Println("After modifySlice():", numbers)

	subSlice := numbers[1:4]

	fmt.Println("Sub Slice:", subSlice)

	fmt.Printf("Length: %d\n", len(subSlice))
	fmt.Printf("Capacity: %d\n", cap(subSlice))

	updated := appendValue(numbers)

	fmt.Println("After append:", updated)

	fmt.Println("Original still accessible:", numbers)
}