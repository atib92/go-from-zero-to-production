# Notes

## Interfaces Describe Behavior

Interfaces focus on:

"What can this type do?"

instead of:

"What is this type?"

This is a major design difference from inheritance-heavy languages.

---

## Implicit Satisfaction

A type satisfies an interface automatically.

There is no:

```java
implements
```

keyword.

If the methods match, the interface is satisfied.

---

## Interfaces Enable Polymorphism

Example:

Human
Dog
Robot

can all satisfy:

Speaker

if they implement:

Speak() string

---

## Interfaces Are Small

Idiomatic Go prefers:

```go
type Reader interface {
	Read([]byte) (int, error)
}
```

over giant interfaces.

Small interfaces are easier to compose and test.

---

## Composition Over Inheritance

Interfaces are one of the primary ways Go achieves abstraction without inheritance.


## Real-World Interface Usage

Interfaces become most valuable when business logic depends on behavior instead of concrete implementations.

Example:

```go
func processOrder(logger Logger)
```

The function can work with:
- ConsoleLogger
- FileLogger
- DatadogLogger

without modification.

This reduces coupling and improves testability.

---

## Interface Composition

Interfaces can embed other interfaces.

Example:

```go
type ReadWriter interface {
    Reader
    Writer
}
```

This creates a new interface by combining smaller interfaces.

The Go standard library uses this pattern extensively.

---

## Prefer Small Interfaces

Idiomatic Go favors small, focused interfaces.

Examples from the standard library:

```go
io.Reader
io.Writer
io.Closer
```

These can be combined as needed instead of creating large interfaces with many methods.