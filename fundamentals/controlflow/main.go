package main

import "fmt"

func grade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 75 {
		return "B"
	} else if score >= 60 {
		return "C"
	}

	return "F"
}

func numberType(num int) string {
	switch {
	case num%2 == 0:
		return "even"
	default:
		return "odd"
	}
}

func sumUntil(limit int) int {
	total := 0

	for i := 1; i <= limit; i++ {
		total += i
	}

	return total
}

func findFirstEven(numbers []int) (int, bool) {
	for _, num := range numbers {
		if num%2 == 0 {
			return num, true
		}
	}

	return 0, false
}

func main() {
	fmt.Println("Grade:", grade(82))

	fmt.Println("Number Type:", numberType(11))

	fmt.Println("Sum Until 5:", sumUntil(5))

	numbers := []int{1, 3, 5, 8, 9}

	even, found := findFirstEven(numbers)

	if found {
		fmt.Println("First Even:", even)
	} else {
		fmt.Println("No even number found")
	}
}
