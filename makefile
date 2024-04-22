.PHONY: default run build test docs clean

# Variables
APP_NAME=go-categories-api

# Tasks
default: run-with-docs

run:
	@air run
run-with-docs:
	@swag init
	@go run ./cmd/api/main.go
build:
	@go build -o $(APP_NAME) ./cmd/api/main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs