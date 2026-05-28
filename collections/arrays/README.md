# Arrays in Go

Arrays in Go are fixed-size collections of elements of the same type.

Example:

```go
var numbers [5]int
```

This creates an array of exactly 5 integers.

## Important Characteristics

- fixed size
- contiguous memory
- value semantics
- size is part of type

## Important Insight

In Go:
```go
[5]int
```

and:

```go
[10]int
```

are different types.

## Arrays vs Slices

Arrays:
- fixed size
- copied by value

Slices:
- dynamic
- lightweight views over arrays

Slices are built on top of arrays.