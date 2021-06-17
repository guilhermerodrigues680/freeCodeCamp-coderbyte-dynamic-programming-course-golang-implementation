## WebAssembly
WebAssembly · golang/go Wiki [github.com/golang/go/wiki/WebAssembly](https://github.com/golang/go/wiki/WebAssembly)

```console
GOOS=js GOARCH=wasm go build -o main.wasm
```

## WebAssembly + TinyGo

https://tinygo.org/getting-started/install/using-docker/
https://levelup.gitconnected.com/a-hands-on-introduction-to-webassembly-with-go-959babb58109
https://wasmbyexample.dev/examples/hello-world/hello-world.go.en-us.html

```console
docker run --rm -v $(pwd):/src tinygo/tinygo:0.18.0 tinygo build -o wasm.wasm -target=wasm examples/wasm/export
```