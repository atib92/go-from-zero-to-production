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