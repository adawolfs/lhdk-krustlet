package main

import (
	"syscall/js"

	"github.com/adawolfs/go-wasm/src/function"
)

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("wasmMessage", js.FuncOf(wasmMessage))
	<-done
}

func wasmMessage(this js.Value, args []js.Value) interface{} {
	return function.Message(args[0].String())
}
