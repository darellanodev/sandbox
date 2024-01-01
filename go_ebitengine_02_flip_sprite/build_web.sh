env GOOS=js GOARCH=wasm go build -o ./build_web/embedtest.wasm github.com/darellanodev/go-test-embed
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./build_web