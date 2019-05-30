// Package app contains the implementation
package app

import (
	"context"
	"log"
)

func (c *Server) CreateDocument(ctx context.Context, in *Document) (*Document, error) {
	log.Printf("Received document: %v", in)
	return &Document{Id: 99, Name: in.Name}, nil
}
