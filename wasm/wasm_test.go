//go:build js && wasm

package main

import (
	"flag"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	update = flag.Bool("update", false, "update the golden files of this test")
)

var testDataDir = "testdata"

func goldenValue(t *testing.T, goldenFile string, actual string, update bool) string {
	t.Helper()
	goldenPath := filepath.Join(".", testDataDir, goldenFile+".golden")

	f, err := os.OpenFile(goldenPath, os.O_RDWR|os.O_CREATE, 0666)
	defer f.Close()

	if update {
		_, err := f.WriteString(actual)
		if err != nil {
			t.Fatalf("Error writing to file %s: %s", goldenPath, err)
		}

		return actual
	}

	content, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("Error opening file %s: %s", goldenPath, err)
	}
	return string(content)
}

func TestParse(t *testing.T) {
	// refer: https://github.com/golang/go/wiki/WebAssembly#executing-webassembly-with-nodejs

	// TODO: refactor, this method should also works for ../kfparser/parser_test.go
	// testdata directory should be shared between wasm and kfparser tests
	tests := []struct {
		name   string
		target string
	}{
		{
			name:   "all_features",
			target: "all_features",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kfPath := filepath.Join(".", testDataDir, tt.target+".kf")

			kfDef, err := os.ReadFile(kfPath)
			require.NoError(t, err)

			got, err := parse(string(kfDef))
			require.NoError(t, err)

			want := goldenValue(t, tt.target, got, *update)

			assert.Equal(t, want, got)
		})
	}
}
