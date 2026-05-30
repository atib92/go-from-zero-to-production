# Maps

Maps are Go's built-in hash table implementation.

Example:

```go
scores := map[string]int{}
```

Maps support:

- insert
- lookup
- update
- delete
- iteration

## Important Characteristics

- hash table based
- reference-like semantics
- fast lookups
- dynamic size

## Common Idioms

Existence check:

```go
value, exists := m[key]
```

Set implementation:

```go
visited := map[string]bool{}
```