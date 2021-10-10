all: build test-run

# Pre-requisites
# cp /usr/local/go/misc/wasm/wasm_exec.js assets/

build:
	# Building wasm script...
	@cd cmd; GOOS=js GOARCH=wasm go build -o  ../assets/main.wasm;
	# Build Completed !

test-run:
	# Spawing test serve...
	go run test/http.go
	# Server shut down !
