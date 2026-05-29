# Pointers in Go

Pointers store memory addresses of variables.

Example:

```go
ptr := &value
```

Pointers enable:
- shared mutation,
- efficient memory usage,
- avoiding unnecessary copies.

## Important Operators

### Address Operator

```go
&value
```

returns memory address.

### Dereference Operator

```go
*ptr
```

accesses value at memory location.

## Important Insight

Go pointers are safer than C/C++:
- no pointer arithmetic
- garbage collected
- simpler semantics