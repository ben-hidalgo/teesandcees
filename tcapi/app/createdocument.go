// Package app contains the implementation
package app

import (
	"context"
	"log"
)

func (s *Server) CreateDocument(ctx context.Context, in *Document) (*Document, error) {
	log.Printf("CreateDocument() received document: %v", in)

	log.Printf("CreateDocument() s.Rc=%v", s.Rc)

	return &Document{Id: 99, Name: in.Name}, nil
}
