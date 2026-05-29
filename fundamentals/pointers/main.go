package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func modifyValue(num int) {
	num = 100
}

func modifyWithPointer(num *int) {
	*num = 100
}

func updateAge(user *User) {
	user.Age = 30
}

func nilPointerExample() {
	var ptr *int

	fmt.Println("Nil Pointer:", ptr)

	if ptr == nil {
		fmt.Println("Pointer is nil")
	}
}

func main() {
	number := 10

	fmt.Println("Original:", number)

	modifyValue(number)

	fmt.Println("After modifyValue():", number)

	modifyWithPointer(&number)

	fmt.Println("After modifyWithPointer():", number)

	ptr := &number

	fmt.Println("Pointer Address:", ptr)

	fmt.Println("Value Through Pointer:", *ptr)

	// An example with struct

	user := User{
		Name: "Atib",
		Age:  25,
	}

	fmt.Println("Before:", user)

	updateAge(&user)

	fmt.Println("After:", user)

	// nill pointer example
	nilPointerExample()
}