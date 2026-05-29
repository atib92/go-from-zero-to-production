# Notes

## Go Is Still Pass-By-Value

Even pointers themselves are passed by value.

What gets copied is:
- the address,
- not the underlying data.

---

## Pointers Enable Shared Mutation

Without pointers:
- function parameters are copied,
- mutations affect only local copies.

Pointers allow functions to modify original data.

---

## Slices Internally Use Pointers

Slices internally contain:
- pointer to array
- length
- capacity

Understanding pointers explains slice behavior.

---

## Structs Often Use Pointers

Large structs are commonly passed using pointers to:
- avoid copies,
- improve performance,
- enable mutations.

---

## Go Pointers Are Simpler Than C

Go intentionally removes:
- pointer arithmetic,
- manual memory management,
- direct memory manipulation.

This keeps pointers safer and easier to reason about.