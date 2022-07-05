//go:build js && wasm

package main

import (
	"fmt"
	// "github.com/gopherjs/gopherjs/js"
	. "github.com/siongui/godom/wasm"
	"syscall/js"
)

var signal = make(chan int)

var foo = Document.GetElementById("foo")
var count = 0

func keepAlive() {
	for {
		<-signal
	}
}

func clicked(this js.Value, args []js.Value) interface{} {
	count++
	foo.Set("textContent", fmt.Sprintf("I am clicked %d time", count))
	return js.ValueOf(count)
}

func main() {
	Window.Alert("hello world!")
	testdivs := Document.QuerySelectorAll("#testdivs > div")
	for _, testdiv := range testdivs {
		testdiv.Set("innerHTML", "hi")
	}
	cb := js.FuncOf(clicked)
	foo.Call("addEventListener", "click", cb)

	keepAlive()
}
