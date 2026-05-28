package main

import "testing"

func TestArrayLength(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	if len(numbers) != 5 {
		t.Errorf("expected length 5")
	}
}

func TestArrayCopy(t *testing.T) {
	original := [3]int{1, 2, 3}

	copied := original

	copied[0] = 100

	if original[0] == 100 {
		t.Errorf("arrays should be copied by value")
	}
}