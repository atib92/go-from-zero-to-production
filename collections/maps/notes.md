# Personal Observations

## Maps Behave Differently From Structs

Struct assignment copies data:

```go
u2 := u1
```

Maps share underlying state:

```go
m2 := m1
```

This makes maps feel much more like slices than structs.

---

## Maps Are Hash Tables

A map stores key/value pairs and provides efficient lookup.

Conceptually:

key
    ↓
hash
    ↓
bucket
    ↓
value

---

## Existence Checks Are Explicit

The pattern:

```go
value, exists := m[key]
```

appears frequently.

This avoids ambiguity between:

- missing key
- zero value

---

## Maps Are Commonly Used As Sets

Go does not have a built-in set type.

Maps are often used instead:

```go
visited := map[string]bool{}
```

---

## Maps Have Shared State

Passing a map into a function can modify the caller's data.

This differs from structs and arrays.

It feels similar to slices.

## A Useful Mental Model
Structs and arrays own data.
Slices and maps point to data.