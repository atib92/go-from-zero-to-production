package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Human struct {
	Name string
}

func (h Human) Speak() string {
	return "Hello, my name is " + h.Name
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! I am " + d.Name
}

// announce doesn't care what the actual type of the object is 
// as long as it implements Speak() i,e to say its a Speaker
func announce(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	human := Human{Name: "Atib"}
	dog := Dog{Name: "Buddy"}
	
	//fmt.Println(dog.Speak())
	//fmt.Println(human.Speak())

	announce(dog)
	announce(human)

	fmt.Println("\n=== Logger Example ===")
	loggerExample()

	fmt.Println("\n=== Interface Composition Example ===")
	compositionExample()
}