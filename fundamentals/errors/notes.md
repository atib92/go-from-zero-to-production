# Personal Observations

## Errors Build On Interfaces

The built-in error type is actually an interface.

type error interface {
    Error() string
}

This was the first place where interfaces felt truly useful.

Instead of special language constructs, Go models errors using ordinary interfaces.

---

## Error Handling Is Explicit

Unlike Python or Java, Go does not primarily rely on exceptions.

The caller explicitly decides what to do with failures.

This makes control flow more visible and predictable.

---

## Errors Are Part Of Function Signatures

Functions often advertise failure possibilities directly:

func findUser(id int) (User, error)

The possibility of failure is visible at the call site.