# Slices in Go

Slices are dynamic, flexible views over arrays.

Example:

```go
numbers := []int{1, 2, 3}
```

Slices are one of the most important data structures in Go.

## Important Characteristics

- dynamic size
- lightweight descriptors
- shared backing arrays
- reference-like behavior

## Slice Internals

A slice internally contains:
- pointer to array
- length
- capacity

## Important Insight

Slices themselves are small values.

But they reference shared underlying memory.

This explains:
- shared mutations
- append behavior
- slicing behavior
- capacity growth