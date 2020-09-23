.POSIX:

.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: fmt
fmt: ## Run all formatings
	go mod vendor
	go mod tidy
	go fmt ./...

.PHONY: run
run: ## Run all examples
	go run ./examples/info/main.go
	@echo ""
	-go run ./examples/warning/main.go
	@echo ""
	-go run ./examples/success/main.go
	@echo ""
	-go run ./examples/error/main.go
	@echo ""
	go run -ldflags="-s -w -X github.com/erdaltsksn/cui.appVersion=v1.0.0" ./examples/version/main.go version
	@echo ""

.PHONY: test
test: ## Run all tests
	go test ./...

.PHONY: coverage
coverage: ## Show test coverage
	@go test -coverprofile=coverage.out ./... > /dev/null
	go tool cover -func=coverage.out
	@rm coverage.out

.PHONY: godoc
godoc: ## Start local godoc server
	@echo "See Documentation:"
	@echo "    http://localhost:6060/pkg/github.com/erdaltsksn/cui"
	@echo ""
	@godoc -http=:6060

.PHONY: clean
clean: ## Clean all generated files
	rm -rf ./vendor/
	rm -rf go.sum
