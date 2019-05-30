FROM golang:1.12.1-stretch

ENV GO111MODULE=on

RUN go get     github.com/golang/protobuf/protoc-gen-go
RUN go install github.com/golang/protobuf/protoc-gen-go

#RUN go get     github.com/fullstorydev/grpcurl
#RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl

RUN apt-get update
RUN apt-get -q -y install protobuf-compiler
RUN apt-get -q -y install net-tools
RUN apt-get -q -y install jq

#
WORKDIR /usr/local/gocache

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download



#WORKDIR /usr/local/greeter_client
#RUN wget https://raw.githubusercontent.com/grpc/grpc-go/master/examples/helloworld/greeter_client/main.go

WORKDIR /usr/local/mount

COPY . .

RUN protoc --go_out=plugins=grpc:. app/tcapi.proto

EXPOSE 8000

#CMD [ "bash", "entrypoint.sh" ]
#CMD [ "sleep", "999999" ]
CMD [ "go", "run", "main.go" ]
