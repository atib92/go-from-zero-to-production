package main

import "testing"

func TestStringLength(t *testing.T) {
	str := "hello"

	if len(str) != 5 {
		t.Errorf("expected 5")
	}
}

func TestRuneCount(t *testing.T) {
	str := "你好"

	if len([]rune(str)) != 2 {
		t.Errorf("expected 2 runes")
	}
}