package usecase

import (
	"context"
	"net/http"
)

// Service ...
type Service struct {
	p2pClient    *http.Client
	hostedClient *http.Client
}

// NewService ...
func NewService(p2pClient, hostedClient *http.Client) *Service {
	return &Service{
		p2pClient:    p2pClient,
		hostedClient: hostedClient,
	}
}

// Fetch ...
func (s *Service) Fetch(ctx context.Context) ([]map[string]interface{}, error) {
	panic("waiting for implementation")
}

// Store ...
func (s *Service) Store(ctx context.Context, items []map[string]interface{}) error {
	panic("waiting for implementation")
	return nil
}
