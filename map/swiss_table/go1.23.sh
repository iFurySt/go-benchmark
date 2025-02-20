#!/bin/sh

docker run -it --rm -v "${PWD}":/go/src/go-benchmark --workdir=/go/src/go-benchmark golang:1.23-bullseye go run map/swiss_table/main.go
