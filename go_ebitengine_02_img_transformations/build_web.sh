env GOOS=js GOARCH=wasm go build -o ./build_web/imgtransformations.wasm github.com/darellanodev/sandbox-go-ebitengine-02-img-transformations
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./build_web