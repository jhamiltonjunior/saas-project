#!/bin/bash

set -e

# Run unit tests
go test -v ./...

# Run integration tests
# TODO: Implement integration tests