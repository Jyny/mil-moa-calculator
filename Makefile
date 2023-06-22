.PHONY: all build clean serve

all: build

serve: build
	go run server/server.go go -addr=0.0.0.0:8080 -root=build

build: clean build/app.wasm build/wasm_exec.js build/index.html 

build/app.wasm: app/app.go
	GOOS=js GOARCH=wasm go build -o build/app.wasm app/app.go

build/wasm_exec.js:
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" build/

build/index.html:
	cp src/index.html build/

clean:
	rm -rf build