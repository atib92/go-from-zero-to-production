package main

import "testing"

func TestDivide(t *testing.T) {
	result, err := divide(10, 2)

	if err != nil {
		t.Fatal(err)
	}

	if result != 5 {
		t.Errorf("expected 5 got %v", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := divide(10, 0)

	if err == nil {
		t.Errorf("expected error")
	}
}