# Deep Dive: Producer / Consumer Pattern

Consider the following implementation:

```go
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
	ch := make(chan int)

	go producer(ch)

	consumer(ch)
}
```

---

## Step 1: Create The Channel

```go
ch := make(chan int)
```

This creates an **unbuffered channel**.

Conceptually:

```text
Channel
    ↓
Empty
```

Because the channel is unbuffered:

```text
Sender blocks until receiver exists.
Receiver blocks until sender exists.
```

The sender and receiver must rendezvous to exchange data.

---

## Why Must One Side Be A Goroutine?

A common question is:

```go
producer(ch)
consumer(ch)
```

Why doesn't this work?

Let's analyze it.

Execution starts with:

```go
producer(ch)
```

The producer enters:

```go
ch <- 1
```

Since this is an unbuffered channel, Go requires a receiver.

However:

```text
consumer() has not started yet
```

because execution is still inside producer().

The producer blocks forever.

The consumer never starts.

The program eventually fails with:

```text
fatal error: all goroutines are asleep - deadlock!
```

---

## The Solution

We start the producer as a goroutine:

```go
go producer(ch)

consumer(ch)
```

Now two execution contexts exist:

```text
Main Goroutine
    ↓
consumer(ch)

Producer Goroutine
    ↓
producer(ch)
```

The producer and consumer can execute concurrently.

This allows:

```text
Producer sends
Consumer receives
```

to occur successfully.

---

## What Happens If Consumer Runs First?

Suppose the scheduler chooses:

```text
consumer()
```

before:

```text
producer()
```

The consumer executes:

```go
value := <-ch
```

No data is available.

Therefore:

```text
Consumer blocks
Waiting for data
```

This is normal.

The scheduler then runs the producer.

The producer executes:

```go
ch <- 1
```

Now:

```text
Producer wants to send
Consumer wants to receive
```

The channel transfers the value.

The consumer wakes up and continues.

---

## What Happens If Producer Runs First?

Suppose the scheduler chooses:

```text
producer()
```

first.

The producer reaches:

```go
ch <- 1
```

No receiver is waiting.

Therefore:

```text
Producer blocks
Waiting for receiver
```

The scheduler then runs the consumer.

The consumer receives.

The value is transferred.

The producer continues.

---

## Important Observation

It does not matter which side arrives first.

The channel synchronizes them.

This is one of the most important properties of channels.

---

# When Is The Channel Closed?

Inside producer:

```go
defer close(ch)
```

This means:

```text
When producer returns,
close the channel.
```

Execution:

```text
send 1
send 2
send 3
send 4
send 5
producer returns
close(ch)
```

The channel is closed after the final value is sent.

---

# What Does close(ch) Mean?

Many beginners assume:

```text
close(ch)
```

means:

```text
Destroy channel
Free memory
```

It does not.

Instead it means:

```text
No more values will be sent.
```

Think of it as:

```text
End Of Stream
```

rather than:

```text
Delete Object
```

---

# Why Do We Need close(ch)?

The consumer uses:

```go
for value := range ch
```

This means:

```text
Receive values
until channel closes.
```

Conceptually:

```go
for {
	value, ok := <-ch

	if !ok {
		break
	}

	fmt.Println(value)
}
```

Without:

```go
close(ch)
```

the consumer would:

```text
Receive 1
Receive 2
Receive 3
Receive 4
Receive 5
Wait forever for value 6
```

Eventually:

```text
deadlock
```

would occur.

---

# Receiving From A Closed Channel

Receiving from a channel actually returns:

```go
value, ok := <-ch
```

Examples:

Channel has data:

```text
value=5
ok=true
```

Channel is closed:

```text
value=0
ok=false
```

This is how:

```go
for value := range ch
```

knows when to stop.

---

# Could The Consumer Be The Goroutine Instead?

Yes.

This also works:

```go
go consumer(ch)

producer(ch)
```

Now:

```text
Main Goroutine
    ↓
producer()

Consumer Goroutine
    ↓
consumer()
```

There is still:

```text
One Sender
One Receiver
```

running concurrently.

The important requirement is:

```text
Sender and Receiver must exist concurrently.
```

Not:

```text
Producer must always be the goroutine.
```

---

# Could Both Be Goroutines?

Absolutely.

Example:

```go
go producer(ch)
go consumer(ch)
```

Now:

```text
Producer Goroutine
Consumer Goroutine
```

both run independently.

However there is a new problem:

```text
What keeps main alive?
```

If main exits:

```text
Process exits
All goroutines terminate
```

Therefore we must synchronize.

Typically:

```go
var wg sync.WaitGroup
```

Example:

```go
var wg sync.WaitGroup

wg.Add(2)

go func() {
	defer wg.Done()
	producer(ch)
}()

go func() {
	defer wg.Done()
	consumer(ch)
}()

wg.Wait()
```

Now:

```text
main waits
until producer completes
until consumer completes
```

before exiting.

---

# The Deep Insight

A channel is much more than a queue.

A channel provides:

1. Data transfer
2. Synchronization

When a sender executes:

```go
ch <- value
```

and a receiver executes:

```go
value := <-ch
```

the channel coordinates both goroutines.

This allows communication without explicit locks.

---

# Mental Model

For an unbuffered channel:

```text
Producer
    ↓
Send
    ↓
Channel
    ↓
Receive
    ↓
Consumer
```

The send and receive must rendezvous.

A useful analogy is:

```text
Two people shaking hands
while passing an object.
```

Neither side can complete the exchange without the other.
