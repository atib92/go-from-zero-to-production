package main

import (
	"testing"
)

func TestBasicChannel(t *testing.T) {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	value := <-ch

	if value != 42 {
		t.Errorf("expected 42 got %d", value)
	}
}

func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 2)

	ch <- 10
	ch <- 20

	first := <-ch
	second := <-ch

	if first != 10 {
		t.Errorf("expected 10 got %d", first)
	}

	if second != 20 {
		t.Errorf("expected 20 got %d", second)
	}
}

func TestProducerConsumer(t *testing.T) {
	ch := make(chan int)

	go producer(ch)

	var results []int

	for value := range ch {
		results = append(results, value)
	}

	if len(results) != 5 {
		t.Errorf("expected 5 values got %d", len(results))
	}

	if results[0] != 1 {
		t.Errorf("expected first value to be 1")
	}

	if results[4] != 5 {
		t.Errorf("expected last value to be 5")
	}
}
