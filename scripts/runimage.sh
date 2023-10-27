#!/bin/bash
set -v

./scripts/buildImage.sh

docker run -p 8080:8080 \
    --expose='8080' \
    -e LINKS_PORT='8080' \
    -e LINKS_BACKEND='postgres' \
    -e LINKS_ADMIN_PATH="/frontend" \
    -e LINKS_BACKEND=postgres \
    -e LINKS_PORT='8080' \
    -e LINKS_DB_HOST=localhost \
    -e LINKS_DB_USER=postgres \
    -e LINKS_DB_PASSWORD=postgres \
    -e LINKS_DB_DATABASE=hcflinks \
    -e LINKS_DB_PORT='5432' \
    links
    # -e LINKS_DB_HOST='8080' \
    # -e LINKS_DB_USER='8080' \
    # -e LINKS_DB_PASSWORD='8080' \
