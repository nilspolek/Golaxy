package main

import (
	"wasmTest/router"
)

func main() {

	r := router.New()
	r.Render()

	select {}
}
