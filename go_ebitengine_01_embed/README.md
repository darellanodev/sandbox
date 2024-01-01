# Go ebitengine embed

## description

It shows how to embed fonts (.ttf), sounds (.wav), images (.png), and text files (.txt) to use with go and ebitengine. Its important when building for web because I got errors when loading external files and also this embed mechanism is most efficient.

## build

Run `./build_web.sh` and it will generate embedtest.wasm and wasm_exec.js inside build_web

## run locally

`go run .`
