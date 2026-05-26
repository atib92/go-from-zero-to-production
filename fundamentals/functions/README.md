# Functions in Go

Functions are first-class citizens in Go.

Go supports:
- typed parameters,
- multiple return values,
- named return values,
- variadic functions,
- closures.

## Important Concepts

- explicit typing
- multiple returns
- error-first design
- functions as values

## Common Patterns

### Returning errors

```go
result, err := someFunction()
if err != nil {
    return err
}
```

This is one of the most common Go idioms.