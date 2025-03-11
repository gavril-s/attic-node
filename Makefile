get-build-targets:
	go tool dist list

build-linux:
	./scripts/build_linux.sh

build-windows: 
	./scripts/build_windows.sh

build-darwin:
	./scripts/build_darwin.sh

build: build-linux build-windows build-darwin

run:
	go run ./cmd/main.go

test:
	go test -v -race -count=2 ./...