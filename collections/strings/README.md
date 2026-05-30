# Strings

Strings in Go are immutable sequences of bytes.

Important concepts:

- UTF-8 encoding
- bytes vs runes
- string immutability
- Unicode handling

## Important Insight

len() returns bytes, not characters.

Example:

```go
len("你好") // 6
```

To count Unicode characters:

```go
len([]rune("你好")) // 2
```

## Common Operations

```go
strings.ToUpper()
strings.Contains()
strings.ReplaceAll()
```