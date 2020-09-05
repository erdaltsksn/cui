.POSIX:

.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: ## Get dependencies
	go get -v -t -d ./...

.PHONY: run
run: ## Run all examples
	bash ./scripts/example.sh

.PHONY: fmt
fmt: ## Run all formatings
	go mod vendor
	go mod tidy
	go fmt ./...

.PHONY: test
test: ## Run all test
	go test -v ./...

.PHONY: build
build: ## Build all
	go build -v ./...
