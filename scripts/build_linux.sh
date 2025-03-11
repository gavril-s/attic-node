#!/bin/bash
archs=(amd64 arm64 mipsle)

for arch in ${archs[@]}
do
	env GOOS=linux GOARCH=${arch} go build -o ./builds/attic-node_linux_${arch} ./cmd/main.go
done