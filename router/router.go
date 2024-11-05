package router

import (
	"strings"
	"syscall/js"

	"wasmTest/component"
)

type Router struct {
	Routes map[string]component.Component
}

func New() Router {
	return Router{
		Routes: make(map[string]component.Component),
	}
}

func (r *Router) RegisterRoute(path string, comp component.Component) {
	r.Routes[strings.ToUpper(path)] = comp
}

func (r Router) Render() {
	var (
		ok   bool
		comp component.Component
	)
	elements := js.Global().Get("document").Call("querySelectorAll", "*")
	for i := 0; i < elements.Length(); i++ {
		element := elements.Index(i)
		if comp, ok = r.Routes[element.Get("tagName").String()]; ok {
			println(element.Get("tagName").String())
			element.Set("innerHTML", comp.Render())
		}
	}
}
