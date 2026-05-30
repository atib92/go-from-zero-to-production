# Personal Observations

## Strings Are Immutable

Characters cannot be modified in-place.

To modify a string:

1. Convert to rune slice
2. Modify runes
3. Convert back to string

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

---

## strings.Builder Exists For Performance

Repeated string concatenation can create many allocations.

Builder provides a more efficient way to construct large strings.