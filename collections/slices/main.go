package main

import "fmt"

/*
Not important to return numbers since the slice will never change.
*/
func modifySlice(numbers []int) {
	numbers[0] = 999
}
/*
  If thers is enough capacity, append is going to use the same backing array
  if there isn't enough capacity, append will allocate a new backing array 
  and the slice returned will point to this new array. This is why it is
  important to return the output slice.
*/
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