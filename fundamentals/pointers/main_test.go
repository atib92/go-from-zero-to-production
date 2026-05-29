package main

import "testing"

func TestModifyWithPointer(t *testing.T) {
	number := 10

	modifyWithPointer(&number)

	if number != 100 {
		t.Errorf("expected 100, got %d", number)
	}
}

func TestUpdateAge(t *testing.T) {
	user := User{
		Name: "Test",
		Age:  20,
	}

	updateAge(&user)

	if user.Age != 30 {
		t.Errorf("expected age 30")
	}
}