SHELL := /usr/bin/env bash

BIN_NAME := psx-memcard
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

.PHONY: build test clean run release help

.PHONY: build
build: # Build the application
	go build -o bin/$(BIN_NAME) main.go


.PHONY: test
test:  # Run unit tests
	go test ./internal/... -v

.PHONY: run
run: # Run the application (development)
	@echo "Running Application..."
	@go run ./main.go



.PHONY: clean
clean: # Clean build artifacts
	rm -f bin/$(BIN_NAME)
	go clean -cache -modcache


.PHONY: snapshot
snapshot: # Build using Docker/BuildKit and output artifacts to ./dist
	DOCKER_BUILDKIT=1 docker build \
		--build-arg GOOS=${GOOS} \
		--build-arg GOARCH=${GOARCH} \
		--progress=plain --rm --output type=local,dest=$$(pwd)/dist .

.PHONY: help
help: # Show this help
	@printf "Available targets:\n\n"
	@awk 'BEGIN {FS = ":.*#"} /^[a-zA-Z0-9_-]+:.*#/ {printf "  \033[1;37m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


.PHONY: redo-release
redo-release: # Undo last release (delete git tag and GitHub release)
	CURRENT_TAG=$$(git describe --tags --abbrev=0); \
	echo "Deleting Git Tag $${CURRENT_TAG}"; \
	git tag -d $${CURRENT_TAG}; \
	git push --delete origin $${CURRENT_TAG}; \
	git tag $${CURRENT_TAG}; \
	git push origin $${CURRENT_TAG};
