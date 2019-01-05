#!/usr/bin/env bash

export GOARCH=amd64
export GOOS=linux
go build -o bin/goeureka goeureka/*.go
docker build -t fruits .
export GOARCH=amd64
export GOOS=darwin
