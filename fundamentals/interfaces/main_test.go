package main

import "testing"

func TestHumanSpeak(t *testing.T) {
	h := Human{Name: "Atib"}

	expected := "Hello, my name is Atib"

	if h.Speak() != expected {
		t.Errorf("expected %s", expected)
	}
}

func TestDogSpeak(t *testing.T) {
	d := Dog{Name: "Buddy"}

	expected := "Woof! I am Buddy"

	if d.Speak() != expected {
		t.Errorf("expected %s", expected)
	}
}