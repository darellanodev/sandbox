env GOOS=js GOARCH=wasm go build -o ./build_web/embed.wasm github.com/darellanodev/sandbox-go-ebitengine-01-embed
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./build_web