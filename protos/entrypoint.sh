#!/usr/bin/env bash

# https://vaneyckt.io/posts/safer_bash_scripts_with_set_euxo_pipefail
set -euxo pipefail


protoc --go_out=plugins=grpc:./artifacts/go tcapi.proto

protoc --ruby_out=./artifacts/ruby tcapi.proto

#protoc --ruby_out=./artifacts/ruby --grpc_out=./artifacts/ruby --plugin=protoc-gen-grpc=`which grpc_ruby_plugin` tcapi.proto

#protoc -I protos/ --ruby_out=lib --grpc_out=lib --plugin=protoc-gen-grpc=`which grpc_ruby_plugin` protos/parser.proto

# root@7efbc41340e1:/usr/local/testtcapi# which grpc_tools_ruby_protoc
# /usr/local/bundle/bin/grpc_tools_ruby_protoc

protoc --js_out=./artifacts/js tcapi.proto

cp -rf artifacts ${MOUNT_DIR}

echo success
#sleep 99999
