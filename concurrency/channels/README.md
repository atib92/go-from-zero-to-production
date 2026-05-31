# Channels

Channels are Go's primary mechanism for communication between goroutines.

They allow goroutines to safely exchange data without explicitly sharing memory.

Example:

```go
ch := make(chan string)
```

Sending:

```go
ch <- "hello"
```

Receiving:

```go
msg := <-ch
```

## Topics Covered

- Basic channels
- Blocking behavior
- Buffered channels
- Producer / Consumer pattern
- Channel closing
- Directional channels

## Unbuffered Channels

```go
ch := make(chan int)
```

An unbuffered channel performs a direct handoff between sender and receiver.

The sender blocks until a receiver is ready.

The receiver blocks until a sender provides data.

## Buffered Channels

```go
ch := make(chan int, 5)
```

Buffered channels maintain an internal queue.

Senders can continue until the buffer becomes full.

## Channel Closing

```go
close(ch)
```

Closing a channel signals:

"No more values will be sent."

Receivers can continue draining existing values.

## Directional Channels

Send-only:

```go
func producer(ch chan<- int)
```

Receive-only:

```go
func consumer(ch <-chan int)
```

These provide compile-time safety and make APIs more expressive.

## Go Philosophy

Do not communicate by sharing memory.

Share memory by communicating.