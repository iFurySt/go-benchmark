#!/bin/sh

docker run -it --rm -v "${PWD}":/go/src/go-benchmark --workdir=/go/src/go-benchmark golang:1.24-bullseye go test -bench=. -benchmem ./map/swissmap_test.go