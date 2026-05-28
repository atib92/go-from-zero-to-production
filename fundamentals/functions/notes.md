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

## Closure
```go
// HTTP Middleware example
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received")
		next.ServeHTTP(w, r)
	})
}
```
The clousure above captures 'next' and adds logging. This patter is extremely usefor for:
- logging
- authentication
- authroization
- metrics
- tracing
- rate limiteing
- ... and more

Production Go servers often look like:

```go
handler := loggingMiddleware(
	authMiddleware(
		tracingMiddleware(
			myHandler,
		),
	),
)
```

Each middleware wraps the next.

This creates a processing pipeline.
