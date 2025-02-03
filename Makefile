build:
	GOOS=linux GOARCH=mipsle go build -o ./builds/attic-node ./cmd/main.go

run:
	go run ./cmd/main.go