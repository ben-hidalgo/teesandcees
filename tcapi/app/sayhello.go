// Package app contains the implementation
package app

import (
	"context"
	"log"
)



// SayHello implements app.GreeterServer
func (c *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &HelloReply{Message: "Hello " + in.Name}, nil
}

