#!/bin/bash
protoc --go-grpc_out=require_unimplemented_servers=false:. \
    --plugin="protoc-gen-ts=./frontend/node_modules/.bin/protoc-gen-ts" \
    --ts_opt=esModuleInterop=true \
    --js_out="./frontend/src/generated" \
    --ts_out="./frontend/src/generated" \
    --go_out=. \
    -I=. ./defs/link.proto

# docker run -v ./defs:/defs namely/protoc-all -f link.proto -l go
# docker run --rm -v ./defs:/defs -v ./lib/gencode:/out -w /defs znly/protoc --python_out=/out -I. link.proto
