# Notes

## Why Multiple Return Values Matter

Go avoids exceptions for most normal error handling.

Instead:

- functions return values explicitly,
- callers decide how to handle failures.

This makes control flow easier to follow in large systems.

## Common Go Idiom

if err != nil

This pattern appears everywhere in production Go.

## Important Insight

In Go:
- errors are not special,
- errors are values.

This is a major philosophy difference from Java/Python/C++.