#!/bin/sh
export GOBIN=$PWD/bin
export GOPATH=$PWD
go build -v -x gen
mv gen bin/gen
