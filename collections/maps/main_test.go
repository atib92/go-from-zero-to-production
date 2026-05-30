package main

import "testing"

func TestMapUpdate(t *testing.T) {
	m := map[string]int{
		"views": 10,
	}

	updateMap(m)

	if m["views"] != 11 {
		t.Errorf("expected 11")
	}
}

func TestMapAssignmentSharesState(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
	}

	m2 := m1

	m2["a"] = 100

	if m1["a"] != 100 {
		t.Errorf("maps should share underlying state")
	}
}