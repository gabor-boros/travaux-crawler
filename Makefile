.PHONY: help build windows linux darwin deps format lint test clean run.dev.windows run.dev.linux run.dev.darwin
.DEFAULT_GOAL := build

BIN_NAME := travaux-crawler

help: ## Show available targets
	@echo "Available targets:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

prerequisites: ## Download and install prerequisites
	go install github.com/goreleaser/goreleaser@latest
	go install github.com/sqs/goreturns@latest

deps: ## Download the dependencies
	go mod download
	go mod tidy

build: deps## Build binary
	goreleaser build --rm-dist --snapshot --single-target
	@find bin -name "$(BIN_NAME)" -exec cp "{}" bin/ \;

release: ## Release a new version on GitHub
	goreleaser release --rm-dist --auto-snapshot

format: deps ## Run formatter on the project
	goreturns -b -local -p -w -e -l .

clean: ## Clean up project root
	rm -rf bin/
