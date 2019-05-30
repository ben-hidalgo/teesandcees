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

	server := &app.Server{}

	err := app.InitRedis(server)
	if err != nil {
		log.Fatal(err)
	}

	address := os.Getenv("LISTEN_ADDRESS")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	app.RegisterTcapiServer(s, server)

	log.Printf("listening on %s", address)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
