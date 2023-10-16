#!/bin/bash
set -v
export LINKS_PORT="8080"
export LINKS_BACKEND="memory"
export LINKS_ADMIN_PATH="./links-admin/out
go get .
go run .