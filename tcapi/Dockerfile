FROM golang:1.12.1-stretch

ENV GO111MODULE=on

RUN go get -u github.com/golang/protobuf/protoc-gen-go

RUN apt-get update
RUN apt-get -q -y install protobuf-compiler
RUN apt-get -q -y install net-tools

#
WORKDIR /usr/local/gocache

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download


WORKDIR /usr/local/mount

#COPY go.mod           go.mod
#COPY go.sum           go.sum
#COPY entrypoint.sh    entrypoint.sh
#COPY main.go          main.go
#COPY app/             app/
#COPY restful/         restful/

EXPOSE 8000

#CMD [ "bash", "entrypoint.sh" ]
CMD [ "sleep", "999999" ]
