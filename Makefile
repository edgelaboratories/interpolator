all: build test

.PHONY: build
build:
	GO111MODULE=on go build -race ./...

.PHONY: test
test:
	GO111MODULE=on go test -race ./...
