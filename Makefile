.PHONY: build
build: go-build

.PHONY: clean
clean: go-clean

.PHONY: test
test: go-test

.PHONY: test
run-local: go-run-local

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

go-build:
	@echo "Building Go services..."
	@rm -rf build
	@mkdir build
	go build -o build -v ./...
	@echo "Go services available at ./build"

go-clean: go-clean-cache go-clean-deps

go-clean-cache:
	@echo "Cleaning build cache..."
	go clean -cache

go-clean-deps:
	@echo "Cleaning dependencies..."
	go mod tidy

go-clean-test-cache:
	@echo "Cleaning test cache..."
	go clean -testcache

go-test:
	@echo "Running tests..."
	go test -v

go-run-local:
	go run cmd/dma-watcher/main.go

.DEFAULT_GOAL := help