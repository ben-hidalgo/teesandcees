#!/usr/bin/env bash

# https://vaneyckt.io/posts/safer_bash_scripts_with_set_euxo_pipefail
set -euxo pipefail

grpc_tools_ruby_protoc -I ./protos --ruby_out=lib --grpc_out=lib ./protos/tcapi.proto

cp lib/tcapi_pb.rb ${MOUNT_DIR}/lib/tcapi_pb.rb

cp lib/tcapi_services_pb.rb ${MOUNT_DIR}/lib/tcapi_services_pb.rb

sleep ${SLEEP}

ruby run_all.rb

echo success
