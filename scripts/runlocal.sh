#!/bin/bash
set -v
export LINKS_PORT="8080"
export LINKS_BACKEND="memory"
go get .
go run .