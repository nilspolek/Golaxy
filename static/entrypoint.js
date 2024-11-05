const go = new Go();
WebAssembly.instantiateStreaming(fetch("bin/lib.wasm"), go.importObject).then(
  (result) => {
    go.run(result.instance);
  },
);
