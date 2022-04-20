#!/bin/sh

set -eux

rm -f go.sum
go mod tidy
go test ./...
sed -i '4,$d' go.mod
