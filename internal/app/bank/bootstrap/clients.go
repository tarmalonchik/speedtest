package bootstrap

import (
	"context"

	"github.com/tarmalonchik/speedtest/internal/app/bank/config"
	"github.com/tarmalonchik/speedtest/internal/pkg/inmemory"
)

type ClientsContainer struct {
	cache *inmemory.InMemory
}

func getClients(ctx context.Context, conf *config.Config) (*ClientsContainer, error) {
	clients := &ClientsContainer{
		cache: inmemory.New(),
	}
	return clients, nil
}
