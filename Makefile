PROGRAM_NAME = find-course

.PHONY: help clean dep lint test build build-docker

.DEFAULT_GOAL := help

help: ## Display this help screen.
	@echo "Makefile available targets:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  * \033[36m%-15s\033[0m %s\n", $$1, $$2}'

clean: ## Clean build directory.
	rm -f ./bin/${PROGRAM_NAME}
	rmdir ./bin

dep: ## Download the dependencies.
	go get -v -t -d ./...

lint: dep ## Lint the source files
	golangci-lint run --timeout 5m -E golint -e '(struct field|type|method|func) [a-zA-Z`]+ should be [a-zA-Z`]+'

test: dep ## Run tests
	go test -v ./...

build: dep ## Build executable.
	go build -v ./...

build-docker: dep ## Build docker.
	go build -o find-course ./cmd/app