# Control Flow in Go

Go emphasizes simple and explicit control flow.

Key concepts:
- if/else
- switch
- for loops
- range loops
- early returns

## Important Philosophy

Go prefers:
- readability,
- explicit branching,
- shallow nesting,
- predictable flow.

## Important Features

### Only One Loop

Go has only:

```go
for
```

This simplifies the language significantly.

### Switch Statements

Go switch statements:
- automatically break,
- can evaluate expressions,
- can replace many chained if statements.

### Range Loops

Range loops are heavily used for:
- slices,
- maps,
- channels,
- strings.