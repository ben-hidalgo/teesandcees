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

	key := uuid.String()

	//var m map[string]interface{}

	m := make(map[string]interface{})

	m["name"]= in.Name

	xx := s.Rc.HMSet(key, m)

	log.Printf("CreateDocument() xx=%#v", xx)

	return &Document{Uuid: key, Name: in.Name}, nil
}
