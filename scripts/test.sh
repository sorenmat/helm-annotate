#!/bin/bash -ex
export CGO_ENABLED=0
go build
go test ./...
