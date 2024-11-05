build:
	GOOS=js GOARCH=wasm go build -o bin/lib.wasm lib.go
	go build -o bin/server server/server.go

run: build
	bin/server

