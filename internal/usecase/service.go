package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// Service ...
type Service struct {
	fetchURL string
	storeURL string
	client   *http.Client
}

// NewService ...
func NewService(client *http.Client, fetchURL, storeURL string) *Service {
	return &Service{
		client:   client,
		fetchURL: fetchURL,
		storeURL: storeURL,
	}
}

// Fetch is a fetch method for implementation
func (s *Service) Fetch(ctx context.Context) ([]map[string]interface{}, error) {
	panic("waiting for implementation")
}

// Store is a store method for implementation
func (s *Service) Store(ctx context.Context, items []map[string]interface{}) error {
	panic("waiting for implementation")
	return nil
}
