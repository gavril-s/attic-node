#!/bin/bash
archs=(amd64 arm64)

for arch in ${archs[@]}
do
	env GOOS=darwin GOARCH=${arch} go build -o ./builds/attic-node_darwin_${arch} ./cmd/main.go
done