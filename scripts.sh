#!/bin/bash

coverage() {
    echo "start";
    go test ./... -coverprofile=coverage.out;
    go tool cover -html=coverage.out;
    echo "end";
}

"$@"
