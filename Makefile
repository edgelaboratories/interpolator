all: build test

.PHONY: build
build:
	go build -race ./...

.PHONY: test
test:
	go test -race ./...
