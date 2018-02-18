#!/bin/bash

echo "Building Windows"
GOOS=windows GOARCH=amd64 go build -v -a

echo "Building Mac"
GOOS=darwin GOARCH=amd64 go build -v -a

echo "Building Linux"
GOOS=linux GOARCH=amd64 go build -v -a -o pinga-linux
