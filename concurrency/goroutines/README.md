# Goroutines

Goroutines are lightweight concurrent functions.

A normal function call:

```go
doWork()
```

A goroutine:

```go
go doWork()
```

The `go` keyword instructs the Go runtime scheduler to execute the function concurrently.

## Topics Covered

- Basic goroutines
- Multiple concurrent goroutines
- WaitGroups
- Closure capture pitfalls

## Important Insight

When the main goroutine exits, the entire process exits. All running goroutines are terminated.