#!/bin/bash
set -v
docker build -t links . 
docker run -p 8080:8080 \
    --expose='8080' \
    -e LINKS_PORT='8080' \
    -e LINKS_BACKEND='memory' \
    -e LINKS_ADMIN_PATH="/frontend" \
    links
    # -e LINKS_DB_HOST='8080' \
    # -e LINKS_DB_USER='8080' \
    # -e LINKS_DB_PASSWORD='8080' \



 links 