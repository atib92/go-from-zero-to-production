package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func updateMap(m map[string]int) {
	m["views"]++
}

func main() {

	// Creation
	scores := map[string]int{
		"alice": 95,
		"bob":   88,
	}

	fmt.Println(scores)

	// Insert
	scores["charlie"] = 92

	// Read
	fmt.Println("Alice:", scores["alice"])

	// Update
	scores["alice"] = 100

	// Delete
	delete(scores, "bob")

	fmt.Println(scores)

	// Existence Check
	score, exists := scores["bob"]

	fmt.Println("Score:", score)
	fmt.Println("Exists:", exists)

	// Iteration
	fmt.Println("\nIterating Map")

	for name, score := range scores {
		fmt.Printf("%s -> %d\n", name, score)
	}

	// Map of Structs
	users := map[string]User{
		"u1": {
			Name: "Atib",
			Age:  35,
		},
	}

	fmt.Println(users["u1"])

	// Reference Semantics Example
	metrics := map[string]int{
		"views": 100,
	}

	updateMap(metrics)

	fmt.Println(metrics)
}