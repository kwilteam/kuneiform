.DEFAULT_GOAL := help

.PHONY: help download wasm build

help:
	@# 20s is the width of the first column
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

download: ## download dependencies
	@echo Download go.mod dependencies
	@go mod download

build: ## build from current commit
	@rm -f ./wasm/*.wasm
	@make antlr
	@make wasm
	@echo Build parser
	@goreleaser build --snapshot --clean
	@rm -f ./sql_grammar/go.mod

wasm: ## build wasm binary
	@echo Build wasm binary
	@go generate ./wasm/wasm.go

antlr: ## build Antlr code
	@echo Generate antlr code for sql-grammar
	@rm -f ./sqlgrammar/*.{go,interp,tokens}
	@cd ./sqlgrammar/ && ./generate.sh

	@echo Generate antlr code for kuneiform-grammar
	@rm -f ./kfgrammar/*.{go,interp,tokens}
	@cd ./kfgrammar/ && ./generate.sh

git-sync: ## sync submodule
	@git submodule update --remote

git-init: ## init submodule
	@git submodule update --init

release: ## release
	@# need configure github token
	@# either `GITHUB_TOKEN` environment or ~/.config/goreleaser/github_token
	@goreleaser release