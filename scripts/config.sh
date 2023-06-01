#!/usr/bin/env sh

# tell Go to avoid using the public servers to fetch the code
go env -w GOPRIVATE=github.com/kwilteam/kwil-db
# tell git to use ssh instead of https
echo '[url "ssh://git@github.com/kwilteam/"]\n\tinsteadOf = https://github.com/kwilteam/' >> ~/.gitconfig