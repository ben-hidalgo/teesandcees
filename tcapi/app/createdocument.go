// Package app contains the implementation
package app

import (
	"context"
	"github.com/google/uuid"
	"log"
)

func (s *Server) CreateDocument(ctx context.Context, in *Document) (*Document, error) {
	log.Printf("CreateDocument() received document: %v", in)

	log.Printf("CreateDocument() s.Rc=%v", s.Rc)

	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &Document{Uuid: uuid.String(), Name: in.Name}, nil
}
