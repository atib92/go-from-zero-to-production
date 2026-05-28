# Notes

## Slices Are Views Over Arrays

Slices do not store data themselves.

They reference underlying arrays.

This is one of the most important Go concepts.

---

## Slicing Does Not Copy

Example:

sub := numbers[1:4]

creates another view into the same array.

Mutations may affect both slices.

---

## Length vs Capacity

Length:
- accessible elements

Capacity:
- available backing array space

Capacity becomes very important for performance optimization.

---

## Append May Allocate

append() may:
- reuse existing backing array,
- or allocate a new one.

This behavior affects:
- performance,
- memory usage,
- concurrency safety.

---

## Slices Are Used Everywhere

Slices are fundamental in Go:
- APIs
- JSON
- networking
- databases
- concurrency
- streaming systems

Understanding slices deeply is critical for becoming effective in Go.