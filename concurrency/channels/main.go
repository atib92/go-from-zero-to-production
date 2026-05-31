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
	// defer closing the channel to signal that no more values will be sent
	// Can we still read from this channel ? Yes, we can read from a closed channel until it's empty.
	// Once it's empty, further reads will return the zero value of the channel's type.
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

	// why is the producer a goroutine ? Because the producer is sending values to the channel,
	// and if it were not a goroutine, it would block the main function until it finishes sending all values.
	// By making it a goroutine, we allow the main function to continue executing and start consuming values from the
	// channel concurrently.

	go producer(ch)

	// Does the consumer block the main function ?
	// Yes, the consumer blocks the main function because it is waiting for values to be sent to the channel.
	// However, since the producer is running in a separate goroutine,
	// it can send values to the channel without blocking the main function.

	// What gurantees that the consumer will receive all values sent by the producer ?
	// The channel itself guarantees that all values sent by the producer will be received by the consumer,
	// as long as the channel is not closed before the consumer starts receiving values.
	// In this example, the producer sends 5 values to the channel and then closes it,
	// allowing the consumer to receive all 5 values without any issues.

	// What happens if the go scheduler decides to run the producer completely before the consumer starts receiving values ?
	// If the go scheduler decides to run the producer completely before the consumer starts receiving values,
	// the producer will send all values to the channel and then close it. The consumer will then start receiving values from the channel,
	// and since the channel is closed, it will receive all values until the channel is empty, at which point it will stop receiving values.
	// This is a common pattern in Go, where the producer sends values to a channel and then closes it to signal that no more values will be sent.
	// The consumer can then range over the channel to receive all values until the channel is empty, without worrying about synchronization issues.

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
