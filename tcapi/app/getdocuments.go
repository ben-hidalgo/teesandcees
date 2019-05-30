// Package app contains the implementation
package app

import (
	"context"
	"log"
)

func (c *Server) GetDocuments(ctx context.Context, in *DocumentQuery) (*DocumentList, error) {
	log.Printf("Received document: %v", in)

	doc1 := &Document{Id: 1, Name: "d1"}

	doc2 := &Document{Id: 2, Name: "d2"}

	docs := []*Document{doc1, doc2}

	return &DocumentList{Documents: docs}, nil
}
