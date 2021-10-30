#!/bin/bash

golangci-lint run || { echo 'lint failed' ; exit 1; }

rm ./build/*
env GOOS=windows GOARCH=amd64 go build -o ./build/snake.windows.amd64.exe .
env GOOS=linux GOARCH=amd64 go build -o ./build/snake.linux.amd64 .
env GOOS=linux GOARCH=arm64 go build -o ./build/snake.linux.arm64 .
env GOOS=darwin GOARCH=amd64 go build -o ./build/snake.darwin.amd64 .