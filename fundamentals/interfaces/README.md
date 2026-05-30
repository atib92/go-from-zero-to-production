# Interfaces

Interfaces define behavior.

An interface is a collection of method signatures.

Example:

```go
type Speaker interface {
	Speak() string
}
```

Any type implementing:

```go
Speak() string
```

automatically satisfies the interface.

## Important Characteristics

- Implicit satisfaction
- Behavior-based abstraction
- Composition-friendly
- No inheritance

## Philosophy

Program against capabilities, not concrete implementations.