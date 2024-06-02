#!/bin/bash

# Change directory to the directory containing the executable
cd "$(dirname "$0")"

# Run the executable with the provided arguments
./tig "$@"
