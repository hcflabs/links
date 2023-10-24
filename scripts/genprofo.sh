#!/bin/bash
protoc --go_out=./lib/gencode --go-grpc_out=require_unimplemented_servers=false:. -l=go --proto_path=. -I=. ./defs/link.proto
# docker run -v ./defs:/defs namely/protoc-all -f link.proto -l go 
# docker run --rm -v ./defs:/defs -v ./lib/gencode:/out -w /defs znly/protoc --python_out=/out -I. link.proto