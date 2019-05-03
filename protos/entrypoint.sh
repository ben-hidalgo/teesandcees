#!/usr/bin/env bash

# https://vaneyckt.io/posts/safer_bash_scripts_with_set_euxo_pipefail
set -euxo pipefail


protoc --go_out=plugins=grpc:./artifacts/go tcapi.proto

protoc --ruby_out=./artifacts/ruby tcapi.proto

protoc --js_out=./artifacts/js tcapi.proto

cp -rf artifacts /usr/local/mount/

sleep 99999
