package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(2, 3)

	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

func TestDivide(t *testing.T) {
	result, err := divide(10, 2)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result != 5 {
		t.Errorf("expected 5, got %f", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := divide(10, 0)

	if err == nil {
		t.Errorf("expected error but got nil")
	}
}