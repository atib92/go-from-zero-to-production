package main

import "testing"

func TestStructCopy(t *testing.T) {
	user1 := User{
		Name: "Atib",
	}

	user2 := user1

	user2.Name = "John"

	if user1.Name == user2.Name {
		t.Errorf("structs should be copied")
	}
}

func TestStructFieldUpdate(t *testing.T) {
	user := User{
		Age: 25,
	}

	user.Age = 30

	if user.Age != 30 {
		t.Errorf("expected age 30")
	}
}

// Value receiver test
func TestGreeting(t *testing.T) {
	user := User{
		Name: "Atib",
	}

	expected := "Hello, my name is Atib"

	if user.Greeting() != expected {
		t.Errorf("expected %s", expected)
	}
}

// Pointer receiver test
func TestBirthday(t *testing.T) {
	user := User{
		Age: 20,
	}

	user.Birthday()

	if user.Age != 21 {
		t.Errorf("expected age 21")
	}
}