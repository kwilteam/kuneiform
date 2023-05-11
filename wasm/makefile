build: wasm.go
	@echo "Build Kuneiform wasm..."
	@GOOS=js GOARCH=wasm go build -o kuneiform.wasm -ldflags "-s -w" -trimpath -tags netgowasm.go
