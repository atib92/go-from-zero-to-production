package main

import "testing"

func TestModifySlice(t *testing.T) {
	numbers := []int{1, 2, 3}

	modifySlice(numbers)

	if numbers[0] != 999 {
		t.Errorf("expected first element to be modified")
	}
}

func TestAppendValue(t *testing.T) {
	numbers := []int{1, 2, 3}

	updated := appendValue(numbers)

	if len(updated) != 4 {
		t.Errorf("expected length 4")
	}
}