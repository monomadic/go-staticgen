#!/bin/sh
export GOBIN=$PWD/bin
export GOPATH=$PWD
go build -v gen
mkdir ./bin
mv ./gen ./bin/gen
