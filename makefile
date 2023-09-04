.DEFAULT_GOAL := help

.PHONY: help download git-sync build-antlr wasm build

help:
	@# 20s is the width of the first column
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

download: ## download dependencies
	@echo Download go.mod dependencies
	@go mod download

git-sync: ## sync submodule
	@git submodule update

build-antlr: ## generate antlr code
	@echo Generate antlr code
	@rm -rf ./kfgrammar/{*.go,*.interp,*.tokens}
	@cd ./kuneiform-grammar && ./generate.sh Go kfgrammar ../kfgrammar

build: ## build from current commit
	@make build-antlr
	@rm -f ./wasm/*.wasm
	@make wasm
	@echo Build parser
	@goreleaser build --snapshot --clean
	@rm -f ./sql_grammar/go.mod

build-go: ## build from current commit using just go
	@rm -f ./wasm/*.wasm
	@make wasm
	@echo Build parser
	@go build -o ./dist/tmp/kuneiform ./cmd/root.go
	@rm -f ./sql_grammar/go.mod

wasm: ## build wasm binary
	@echo Build wasm binary
	@go generate ./wasm/wasm.go

release: ## release
	@# need configure github token
	@# either `GITHUB_TOKEN` environment or ~/.config/goreleaser/github_token
	@goreleaser release

test: ## run tests
	@go test -v ./kfparser/ -count=1
