.DEFAULT_GOAL := help
SHELL := /usr/bin/env bash

install-tools: ## Install some tools needed for linting and releasing
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/goreleaser/goreleaser@latest
	go install github.com/caarlos0/svu@latest

format: ## Run `go fmt` on all .go files
	go fmt ./...

lint: ## Run golangci-lint
	golangci-lint run

run: ## Build and run
	go run .

build: ## Build with goreleaser
	goreleaser build --clean --snapshot --single-target

build-all: ## Build with goreleaser for all architectures
	goreleaser build --clean --snapshot

tag: ## Make an annotated git tag for the next version
	git tag -a $$(svu next) -m $$(svu next)

tag-push: tag ## Tag and push up
	git push --follow-tags

release-dry-run: ## Build binaries and simulate a release locally
	goreleaser release --snapshot --clean

# From: https://www.thapaliya.com/en/writings/well-documented-makefiles/
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
