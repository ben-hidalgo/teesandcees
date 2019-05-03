// Package helloworld contains the implementation
package helloworld

import (
	"context"
	"log"
)



// SayHello implements helloworld.GreeterServer
func (c *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &HelloReply{Message: "Hello " + in.Name}, nil
}

