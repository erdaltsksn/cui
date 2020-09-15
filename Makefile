.POSIX:

.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init: ## Get dependencies
	go get -v -t -d ./...

.PHONY: fmt
fmt: ## Run all formatings
	go mod vendor
	go mod tidy
	go fmt ./...

.PHONY: run
run: ## Run all examples
	go run ./examples/info/main.go
	-go run ./examples/warning/main.go
	-go run ./examples/success/main.go
	-go run ./examples/error/main.go
	go run -ldflags="-s -w -X github.com/erdaltsksn/cui.appVersion=v1.0.0" ./examples/version/main.go version

.PHONY: test
test: ## Run all tests
	go test -v ./...

.PHONY: coverage
coverage: ## Show test coverage
	@go test -coverprofile=coverage.out ./... > /dev/null
	go tool cover -func=coverage.out
	rm coverage.out

.PHONY: godoc
godoc: ## Start local godoc server
	@echo "See Documentation:"
	@echo "\thttp://localhost:6060/pkg/github.com/erdaltsksn/cui"
	@echo "\n"
	@godoc -http=:6060

.PHONY: build
build: ## Build the app
	go build -v ./...

.PHONY: clean
clean: ## Clean all generated files
	rm -rf ./vendor/
