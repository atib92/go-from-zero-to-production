package main

import "testing"

func TestGrade(t *testing.T) {
	result := grade(95)

	if result != "A" {
		t.Errorf("expected A, got %s", result)
	}
}

func TestNumberType(t *testing.T) {
	result := numberType(4)

	if result != "even" {
		t.Errorf("expected even, got %s", result)
	}
}

func TestSumUntil(t *testing.T) {
	result := sumUntil(5)

	if result != 15 {
		t.Errorf("expected 15, got %d", result)
	}
}

func TestFindFirstEven(t *testing.T) {
	numbers := []int{1, 3, 7, 8}

	result, found := findFirstEven(numbers)

	if !found {
		t.Errorf("expected even number to be found")
	}

	if result != 8 {
		t.Errorf("expected 8, got %d", result)
	}
}
