// Package app contains the implementation
package app

import (
	"context"
	"log"
)

func (s *Server) GetDocuments(ctx context.Context, in *DocumentQuery) (*DocumentList, error) {

	log.Printf("GetDocuments() Received document: %v", in)

	log.Printf("GetDocuments() s.Rc=%v", s.Rc)

	doc1 := &Document{Uuid: "document:1", Name: "d1"}

	doc2 := &Document{Uuid: "document:2", Name: "d2"}

	docs := []*Document{doc1, doc2}

	return &DocumentList{Documents: docs}, nil
}
