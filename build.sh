#!/bin/bash
# Change directory to the cmd/tig directory
cd cmd/tig

# Build the Go program into an executable named "tig"
go build -o ../../tig

# Change back to the root directory
cd ../..
