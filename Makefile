.PHONY: test all hal

test:
	go test ./...

all:
	go build ./...

hal:
	go run ./cmd/
