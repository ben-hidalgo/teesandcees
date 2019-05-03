// Package main implements a server for Greeter service.
package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"tcapi/app"
)


func main() {

	address := os.Getenv("LISTEN_ADDRESS")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	app.RegisterTcapiServer(s, &app.Server{})

	log.Printf("listening on %s", address)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
