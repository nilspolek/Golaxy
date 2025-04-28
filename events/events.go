package events

import (
	"errors"
	"syscall/js"
)

var ErrNoElementFound = errors.New("Element not found")

const (
	ON_CLICK     = "onclick"
	ON_CHANGE    = "onchange"
	ON_MOUSEOVER = "onmouseover"
	ON_MOUSEOUT  = "onmouseout"
	ON_KEYDOWN   = "onkeydown"
	ON_LOAD      = "onload"

	SELECTOR_CLASS = "."
	SELECTOR_ID    = "#"
)

func BindEvent(selector, event string, handler func(js.Value, []js.Value) interface{}) error {
	document := js.Global().Get("document")
	element := document.Call("querySelector", selector)
	if !element.Truthy() {
		return ErrNoElementFound
	}
	element.Set(event, handler)
	return nil
}

func BindToId(selector, event string, handler func(js.Value, []js.Value) interface{}) error {
	return BindEvent(SELECTOR_ID+selector, event, handler)
}

func BindToClass(selector, event string, handler func(js.Value, []js.Value) interface{}) error {
	return BindEvent(SELECTOR_CLASS+selector, event, handler)
}
