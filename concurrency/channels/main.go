package main

import (
	"fmt"
	"time"
)

func basicChannel() {
	fmt.Println("\n=== Basic Channel ===")

	ch := make(chan string)

	go func() {
		ch <- "hello from goroutine"
	}()

	msg := <-ch

	fmt.Println(msg)
}

func blockingBehavior() {
	fmt.Println("\n=== Blocking Behavior ===")

	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)

		ch <- 42
	}()

	fmt.Println("Waiting for data...")

	value := <-ch

	fmt.Println("Received:", value)
}

func bufferedChannel() {
	fmt.Println("\n=== Buffered Channel ===")

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func producer(ch chan<- int) {
	defer close(ch)

	for i := 1; i <= 5; i++ {
		ch <- i
	}
}

func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("Consumed:", value)
	}
}

func producerConsumer() {
	fmt.Println("\n=== Producer Consumer ===")

	ch := make(chan int)

	go producer(ch)

	consumer(ch)
}

func directionalChannels() {
	fmt.Println("\n=== Directional Channels ===")

	ch := make(chan int)

	go producer(ch)

	consumer(ch)
}

func main() {
	basicChannel()

	blockingBehavior()

	bufferedChannel()

	producerConsumer()

	directionalChannels()
}
