## WebAssembly
WebAssembly · golang/go Wiki [github.com/golang/go/wiki/WebAssembly](https://github.com/golang/go/wiki/WebAssembly)

```console
GOOS=js GOARCH=wasm go build -o main.wasm
```

## WebAssembly + TinyGo

```console
$ tinygo version
tinygo version 0.18.0 darwin/amd64 (using go version go1.16.3 and LLVM version 11.0.0)
```

https://tinygo.org/getting-started/install/
https://tinygo.org/getting-started/install/using-docker/
https://levelup.gitconnected.com/a-hands-on-introduction-to-webassembly-with-go-959babb58109
https://wasmbyexample.dev/examples/hello-world/hello-world.go.en-us.html

```console
docker run --rm -v $(pwd):/src tinygo/tinygo:0.18.0 tinygo build -o wasm.wasm -target=wasm examples/wasm/export

tinygo build -o main.wasm -target=wasm main.go
```