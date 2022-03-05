NAME := $(notdir $(PWD))

PHONY: deps
deps:
	@go mod tidy

PHONY: fmt
fmt:
	@go fmt

PHONY: lint
lint:
	@go vet

PHONY: build
build:
	@CGO_ENABLED=0 go build -o bin/$(NAME) *.go

PHONY: generate-client
generate-client:
	@oapi-codegen -generate "types" -package jamfproapi openapi.json > types.go
	@oapi-codegen -generate "client" -package jamfproapi openapi.json > client.go
	@$(MAKE) fmt
