package main

import (
	"fmt"
	"sync"
	"time"
)

func basicGoroutine() {
	fmt.Println("\n=== Basic Goroutine ===")

	go func() {
		fmt.Println("Hello from goroutine")
	}()

	time.Sleep(500 * time.Millisecond)
}

func multipleGoroutines() {
	fmt.Println("\n=== Multiple Goroutines ===")

	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Worker: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for ch := 'A'; ch <= 'E'; ch++ {
			fmt.Printf("Letter: %c\n", ch)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
}

func waitGroupExample() {
	fmt.Println("\n=== WaitGroup Example ===")

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 1; i <= 3; i++ {
			fmt.Printf("Task-1: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 1; i <= 3; i++ {
			fmt.Printf("Task-2: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()

	fmt.Println("All goroutines completed")
}

func closureCaptureBug() {
	fmt.Println("\n=== Closure Capture Bug ===")

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			fmt.Printf("BUGGY i=%d\n", i)
		}()
	}

	wg.Wait()
}

func closureCaptureFix() {
	fmt.Println("\n=== Closure Capture Fix ===")

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(num int) {
			defer wg.Done()

			fmt.Printf("FIXED i=%d\n", num)
		}(i)
	}

	wg.Wait()
}

func main() {
	basicGoroutine()

	multipleGoroutines()

	waitGroupExample()

	closureCaptureBug()

	closureCaptureFix()
}
