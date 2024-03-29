#!/usr/bin/env bash

go mod tidy

STATUS=$(git status go.mod go.sum --porcelain)

if [[ -n "$STATUS" ]]; then
    echo "Changes detected in 'go.mod' or 'go.sum' after running 'go mod tidy'!"
    exit 1
else
    echo "No changes detected in 'go.mod' or 'go.sum'."
    exit 0
fi