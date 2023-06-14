//go:build js && wasm

//go:generate env GOOS=js GOARCH=wasm go build -o kuneiform.wasm -ldflags "-s -w" -trimpath -tags netgowasm.go
package main

import (
	"fmt"
	"github.com/kwilteam/kuneiform/kfparser"
	"syscall/js"
)

func parse(input string) (json string, err error) {
	schema, err := kfparser.ParseKF(input, nil, kfparser.Default)
	if err != nil {
		return "", err
	}

	jsonBytes, err := schema.ToJSON()
	if err != nil {
		return "", err
	}

	json = string(jsonBytes)
	return
}

// parseWrapper wraps the parse function to be exposed to the global scope
// returns a map {"json": "json output", "error": "error string"}
func parseWrapper() js.Func {
	parseFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		input := args[0].String()
		result := map[string]any{}

		jsonOut, err := parse(input)
		if err != nil {
			errStr := fmt.Sprintf("parsing failed: %s\n", err)
			result["error"] = errStr
		}
		result["json"] = jsonOut
		return result
	})
	return parseFunc
}

func main() {
	fmt.Println("Load Kuneiform parser...")
	// expose the parse function to the global scope
	js.Global().Set("parseKuneiform", parseWrapper())
	<-make(chan bool)
}
