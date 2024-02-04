#!/bin/bash
set -v
export LINKS_PORT="8080"
export LINKS_BACKEND="postgres"
export LINKS_ADMIN_PORT="8081"
export LINKS_DB_HOST="localhost"
export LINKS_DB_USER="postgres"
export LINKS_DB_PASSWORD="postgres"
export LINKS_DB_DATABASE="hcflinks"
export LINKS_DB_PORT="5432"

go get .
go run .