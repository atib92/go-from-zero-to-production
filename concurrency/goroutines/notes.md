# Personal Observations

## Goroutines Feel Extremely Lightweight

Starting a goroutine is as simple as:

```go
go someFunction()
```

This is significantly simpler than thread creation in many other languages.

---

## Main Is Also A Goroutine

The `main()` function runs inside the main goroutine.

When it exits:

- the process exits
- all goroutines stop

---

## WaitGroup Is The First Synchronization Primitive

Using:

```go
time.Sleep(...)
```

is not reliable.

Using:

```go
wg.Wait()
```

allows the program to explicitly wait for work completion.

---

## Output Ordering Is Non-Deterministic

The scheduler decides when goroutines run.

The order of printed output is not guaranteed.

This is one of the first indicators that concurrent programs require different reasoning than sequential programs.

---

## Closures And Goroutines Require Care

Anonymous functions can capture variables from surrounding scope.

When those functions run concurrently, unexpected behavior can occur if shared variables are modified.

Passing variables explicitly into the goroutine is often clearer and safer.


## Comparison With Python Asyncio

Coming from Python, goroutines initially feel very similar to:
- coroutines
- greenlets
- asyncio tasks

Both are:
- lightweight
- managed by a runtime
- intended for massive concurrency

However, the execution model is different.

---

### Python Asyncio Uses Cooperative Scheduling

Example:

```python
async def worker():
    await asyncio.sleep(1)
```

When a coroutine reaches:

```python
await
```

it voluntarily yields control back to the event loop. The event loop can then schedule another coroutine.

Conceptually:

Coroutine A
    ↓ await
Coroutine B
    ↓ await
Coroutine C

This is called:

```text
Cooperative Scheduling
```

because coroutines explicitly cooperate with the scheduler.

---

### Go Uses Preemptive Scheduling

Example:

```go
go worker()
```

A goroutine does not need to explicitly yield. The Go runtime scheduler can pause a goroutine and run another one.

Conceptually:

Goroutine A
Goroutine B
Goroutine C

The scheduler decides when they run.

This is called:

```text
Preemptive Scheduling
```

---

### No async/await Syntax

In Python, asynchronous code often becomes:

```python
await database_call()
await http_request()
await queue.get()
```

and async behavior tends to propagate through the call stack. This is sometimes described as:

```text
async all the way down
```

In Go, ordinary synchronous code can run inside a goroutine:

```go
go func() {
    resp, err := http.Get(url)
}()
```

No special async syntax is required.

---

### Goroutines Can Run On Multiple OS Threads

A typical Python asyncio application often looks like:

Many Coroutines
        ↓
Single Event Loop
        ↓
Single Thread

Go's runtime scheduler is more sophisticated:

Many Goroutines
        ↓
Go Scheduler
        ↓
Many OS Threads
        ↓
Many CPU Cores

This allows goroutines to utilize multiple CPU cores.

---

### Similar Behavior, Different Mechanisms

The following examples appear similar:

Python:

```python
await asyncio.sleep(1)
```

Go:

```go
time.Sleep(1 * time.Second)
```

In both cases:

Task A sleeps
    ↓
Task B runs

However, the mechanism is different.

Python:

```text
Coroutine voluntarily yields
```

Go:

```text
Runtime scheduler parks the goroutine
and schedules another goroutine
```

---

### Personal Mental Model

A useful intuition is:

```text
Goroutines ≈ Asyncio Coroutines
            + Automatic Scheduling
            + Multiple OS Threads
            + Multiple CPU Cores
            - async/await syntax
```

This is not perfectly accurate, but it provides a helpful starting point when transitioning from Python concurrency to Go concurrency.