package main

import (
	// "github.com/gopherjs/gopherjs/js"
	. "github.com/siongui/godom/wasm"
)

func main() {
	// js.Global.Set("pet", map[string]interface{}{
	// 	"New": New,
	// })
	Window.Alert("hello world!")
}

// type Pet struct {
// 	name string
// }

// func New(name string) *js.Object {
// 	return js.MakeWrapper(&Pet{name})
// }

// func (p *Pet) Name() string {
// 	return p.name
// }

// func (p *Pet) SetName(name string) {
// 	p.name = name
// }
