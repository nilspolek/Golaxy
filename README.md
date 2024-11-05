# Golaxy
Golaxy is an easy to use Webframework like react or Angular but it uses Webassambly and is written in go 
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
        <test></test>
    </body>
</html>
```

## Setup
You must serve an entrypoint like 
```javascript
const go = new Go();
WebAssembly.instantiateStreaming(fetch("bin/lib.wasm"), go.importObject).then(
  (result) => {
    go.run(result.instance);
  },
);
```
And an WASM_EXEC that can be found at 
```bash
"$(go env GOROOT)/lib/wasm/wasm_exec.js"
```
