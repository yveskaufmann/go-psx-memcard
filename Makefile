SHELL := /usr/bin/env bash

BIN_NAME := psx-memcard
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

.PHONY: build test clean run run-service test-integration

# Build the application
build:
	go build -o bin/$(BIN_NAME) main.go

# Run unit tests
test:
	go test ./pkg/... -v

run:
	@echo "Running Application..."
	@go run ./main.go


# Clean build artifacts
clean:
	rm -f bin/$(BIN_NAME)
	go clean -cache -modcache

