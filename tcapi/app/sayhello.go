// Package app contains the implementation
package app

import (
	"context"
	"log"
)

// SayHello implements app.GreeterServer
func (s *Server) SayHello(ctx context.Context, hr *HelloRequest) (*HelloReply, error) {

	log.Printf("SayHello() hello: %v", hr.Name)

	log.Printf("SayHello() s.Rc=%v", s.Rc)

	return &HelloReply{Message: "Hello " + hr.Name}, nil
}
