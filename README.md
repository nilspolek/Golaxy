# Golaxy - A Web Framework in Go with WebAssembly
Golaxy is an easy-to-use web framework similar to React or Angular, but written in Go and powered by WebAssembly (WASM). It allows developers to build client-side applications in Go that can run directly in the browser, using WebAssembly for high-performance execution.

## Features
Go as the primary language: Build your frontend applications using Go instead of JavaScript or TypeScript.
WebAssembly powered: Golaxy utilizes WebAssembly (WASM) for running Go code in the browser, enabling fast and efficient execution.
Simple routing: Just like other frameworks, Golaxy provides a simple routing mechanism for managing pages and views.
Lightweight: Golaxy is minimalistic and focused, with only essential features needed for building frontend applications.
Example Code
Go Code
Here's an example of how you can use Golaxy to define a simple route and render a "Hello World" page:

```go
// needs to be compiled to WASM with
// env GOOS=js GOARCH=wasm go build -o main.wasm main.go
package main

import (
	"github.com/nilspolek/Golaxy/router"
)

type Home struct{}

// The Render method defines the HTML to display for this route.
func (h Home) Render() string {
	return "<h1>Hello, World!</h1>"
}

func main() {
	// Initialize the router
	r := router.New()
	// Register a route named "home" that will render the Home struct
	r.RegisterRoute("home", Home{})
	// Render the content of the registered route
	r.Render()
}
```
## HTML Code
The HTML file loads the WebAssembly binary and sets up the entry point to run your Go code.

```html
<!doctype html>
<html>
    <head>
        <meta charset="utf-8" />
        <title>Wasm with Go</title>
        <script src="./static/wasm_exec.js"></script>
        <script src="./static/entrypoint.js"></script>
    </head>
    <body>
        <home></home>
    </body>
</html>
```
## JavaScript Entrypoint
The JavaScript entrypoint is responsible for loading and executing the WebAssembly binary in the browser.

```javascript
const go = new Go();

// Fetch and instantiate the WebAssembly binary
WebAssembly.instantiateStreaming(fetch("bin/lib.wasm"), go.importObject).then(
  (result) => {
    go.run(result.instance);
  },
);
```
## Setup Instructions
To use Golaxy, you'll need to set up a few things:

Install Go: Make sure you have Go installed on your system. You can download it from the official Go website.

Install WebAssembly (WASM) Tools: You'll need to compile your Go code to WebAssembly. Golaxy uses the WebAssembly build toolchain for Go (GOOS=js and GOARCH=wasm).

Serve WebAssembly: Your Go code must be compiled to WASM. Here's how to build your Go application:

```bash
env GOOS=js GOARCH=wasm go build -o main.wasm main.go
```
Create the necessary HTML and JavaScript files:

And an WASM_EXEC that can be found at 

```bash
"cp $(go env GOROOT)/lib/wasm/wasm_exec.js ."
```

Include the wasm_exec.js script provided by Go (it’s typically located in your Go installation directory under $GOROOT/misc/wasm/wasm_exec.js).
Create an HTML file to load the WebAssembly binary and run the Go code.
Run the Application: After setting everything up, run a local server (like http-server in Node.js or Python’s simple HTTP server) to serve your HTML, JavaScript, and WebAssembly files.

Directory Structure
Your project directory should look something like this:

```css
/golaxy-project
├── bin/
│   └── main.wasm
├── static/
│   ├── wasm_exec.js
│   └── entrypoint.js
├── main.go
└── index.html
```
## Requirements
Go 1.11+ (for WebAssembly support).
A modern web browser that supports WebAssembly.
Contributing
Golaxy is open-source! If you'd like to contribute, feel free to open an issue or submit a pull request. We welcome improvements, bug fixes, and documentation updates.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
