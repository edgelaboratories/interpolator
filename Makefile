all: build test

.PHONY: build
build:
	GO111MODULE=on go build -race ./...

.PHONY: test
test:
	mkdir -p bin
	go test -tags=integration --race -coverprofile=bin/cover.out ./...
	go tool cover -html=bin/cover.out -o bin/cover.html
