#!/usr/bin/env bash

# This script ensures that the generated antlr files are up-to-date.

pushd $(git rev-parse --show-toplevel)  > /dev/null

# Force regenerate of pb Go code.
make build-antlr

STATUS=$(git status kfgrammar --porcelain)

popd > /dev/null

if [[ -n "$STATUS" ]]; then
    echo "Changes detected in kfgrammar !"
    exit 1
else
    echo "No changes detected in the directory."
    exit 0
fi

popd