// Package app contains the implementation
package app

import (
	"context"
	"github.com/google/uuid"
	"log"
)


func (d Document) asMap() map[string]interface{} {

	m := make(map[string]interface{})

	m["uuid"] = d.Uuid
	m["name"]= d.Name

	return m
}

func (s *Server) CreateDocument(ctx context.Context, doc *Document) (*Document, error) {
	log.Printf("CreateDocument() received document: %v", doc)

	log.Printf("CreateDocument() s.Rc=%v", s.Rc)

	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	doc.Uuid = uuid.String()

	m := doc.asMap()

	cmd := s.Rc.HMSet(doc.Uuid, m)

	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	return doc, nil
}
