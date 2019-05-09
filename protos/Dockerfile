FROM golang:1.12.1-stretch

RUN go get     github.com/golang/protobuf/protoc-gen-go
RUN go install github.com/golang/protobuf/protoc-gen-go

RUN apt-get update
RUN apt-get -q -y install protobuf-compiler

#WORKDIR /usr/local/deb
#RUN wget http://archive.ubuntu.com/ubuntu/pool/universe/r/ruby-grpc/ruby-grpc-tools_1.3.2+debian-4build1_all.deb

#
WORKDIR /usr/local/protos

RUN mkdir artifacts
RUN mkdir artifacts/js
RUN mkdir artifacts/ruby
RUN mkdir artifacts/go

COPY entrypoint.sh entrypoint.sh
COPY tcapi.proto   tcapi.proto

CMD [ "bash", "entrypoint.sh" ]
