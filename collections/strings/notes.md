# Personal Observations

## Strings Are Immutable

Characters cannot be modified in-place.

To modify a string:

1. Convert to rune slice
2. Modify runes
3. Convert back to string

```go
runes := []rune("hello")
runes[0] = 'H'
str := string(runes)
```

---

## len() Counts Bytes

This surprised me.

Example:

```go
len("你好")
```

returns:

```text
6
```

because UTF-8 uses 3 bytes per character.

---

## range Understands UTF-8

Iterating with:

```go
for _, r := range str
```

automatically decodes UTF-8 runes.

This is usually the correct way to iterate over user-facing text.

---

## Bytes vs Characters

A string is fundamentally bytes.
Characters are represented by runes.
This distinction becomes important when handling Unicode text.

## Runes
A rune is a
```go
type rune = int32
```
and represents a unicode code point.

## Summary of bytes, charecters, unicode and runes
Strings are UTF-8 encoded bytes.
A rune represents a Unicode code point.
A rune may be encoded using 1-4 bytes in UTF-8.
In our chinese text example, each rune was encoded in 3 bytes.

---

## strings.Builder Exists For Performance

Repeated string concatenation can create many allocations.
```go
var builder strings.Builder
builder.WriteString("Hello ")
builder.WriteString("Go")
fmt.Println(builder.String())
```
Builder provides a more efficient way to construct large strings. Very common in production code.
```go
result := ""
for i := 0; i < 10000; i++ {
	result += "a"
}
```