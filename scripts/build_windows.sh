#!/bin/bash
archs=(amd64 arm64)

for arch in ${archs[@]}
do
	env GOOS=windows GOARCH=${arch} go build -o ./builds/attic-node_windows_${arch}.exe ./cmd/main.go
done