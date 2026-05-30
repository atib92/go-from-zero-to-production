# Errors

Errors are values in Go.

The built-in error type is an interface:

type error interface {
    Error() string
}

Go prefers explicit error handling over exceptions.

Common patterns:

- errors.New()
- fmt.Errorf()
- if err != nil