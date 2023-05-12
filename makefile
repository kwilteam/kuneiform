.DEFAULT_GOAL := help

.PHONY: help download wasm build

help:
	@# 20s is the width of the first column
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

download: ## download dependencies
	@echo Download go.mod dependencies
	@go mod download

build: ## build from current commit
	@echo Build parser
	@goreleaser build --snapshot --clean

wasm: ## build wasm binary
	go generate ./wasm/wasm.go