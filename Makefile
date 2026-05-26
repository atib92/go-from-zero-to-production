run:
	go run fundamentals/variables/main.go

test:
	go test ./...

bench:
	go test -bench=. ./...

race:
	go test -race ./...

fmt:
	gofmt -w .

vet:
	go vet ./...