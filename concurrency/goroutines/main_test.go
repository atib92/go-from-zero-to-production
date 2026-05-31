package main

import (
	"sync"
	"testing"
)

func TestWaitGroupSynchronization(t *testing.T) {
	var wg sync.WaitGroup

	counter := 0

	wg.Add(1)

	go func() {
		defer wg.Done()
		counter++
	}()

	wg.Wait()

	if counter != 1 {
		t.Errorf("expected counter to be 1")
	}
}
