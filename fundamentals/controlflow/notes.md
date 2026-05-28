# Notes

## Early Returns

Go code prefers early returns instead of deep nesting.

Instead of:

if something {
    if anotherThing {
        ...
    }
}

Go often prefers:

if err != nil {
    return err
}

This keeps control flow flatter and easier to read.

---

## Range Loops

Range loops are idiomatic in Go.

Example:

for _, value := range items

The underscore `_` ignores unused variables.

---

## Switch Statements

Go switch statements automatically break.

Unlike C/Java:
- explicit break is usually unnecessary.

---

## Infinite Loops

Infinite loops are written as:

for {
}

This is heavily used in:
- servers,
- workers,
- streaming systems,
- event loops.

---

## Go Philosophy

Go intentionally keeps control flow simple.

The language avoids:
- complex syntax,
- implicit behavior,
- hidden control paths.