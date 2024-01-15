# Go ebitengine flip

## description

It shows how to flip images, how to display the real sprite scale and understand coordinate system

## build

Run `./build_web.sh` and it will generate embedtest.wasm and wasm_exec.js inside build_web

## run locally

Execute `./run.sh` or `go run .` command.

## Customize keybindings.json in VSCode

You can use this settings into VSCode `keybindings.json`:

```json
  {
    "key": "ctrl+r",
    "command": "workbench.action.terminal.sendSequence",
    "args": {
      "text": "./run.sh\u000D"
    },
  },
```
