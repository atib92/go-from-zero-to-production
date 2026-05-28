# Notes

## Arrays Are Values

Arrays in Go are copied during:
- assignment
- function calls

Example:

copied := original

creates a completely new array. ( Unlike in Python where it just a reference / shallow copy)

---

## Arrays Have Fixed Size

The size is part of the type.

Example:

[3]int != [4]int

This is fundamentally different from slices.

---

## Arrays Use Contiguous Memory

Arrays are stored in contiguous memory.

This makes:
- indexing fast,
- memory predictable,
- cache locality good.

---

## Arrays Are Rarely Used Directly

In production Go:
- slices are far more common.

However understanding arrays is essential because:
- slices internally reference arrays.

Understanding arrays makes slice behavior much easier to understand later.

---

## Zero Values

Arrays are automatically initialized.

Example:

var numbers [5]int

produces:

[0 0 0 0 0]