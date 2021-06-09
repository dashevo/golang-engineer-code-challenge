package usecase

import (
	"context"
	"net/http"
)

// Service is a service which should interact with p2p and / or self-hosted services in a network
type Service struct {
	p2pClient        *http.Client
	selfHostedClient *http.Client
}

// NewService returns a new service
func NewService(p2pClient, selfHostedClient *http.Client) *Service {
	return &Service{
		p2pClient:        p2pClient,
		selfHostedClient: selfHostedClient,
	}
}

// Fetch returns a fetched data
func (s *Service) Fetch(ctx context.Context) ([]map[string]interface{}, error) {
	panic("waiting for implementation")
}

// Store stored a passed items in external service(s)
func (s *Service) Store(ctx context.Context, items []map[string]interface{}) error {
	panic("waiting for implementation")
	return nil
}
