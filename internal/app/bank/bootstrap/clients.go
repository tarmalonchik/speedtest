package bootstrap

import (
	"context"

	"github.com/tarmalonchik/speedtest/internal/app/bank/config"
	"github.com/tarmalonchik/speedtest/internal/pkg/nodemanager/measurement"

	"github.com/tarmalonchik/speedtest/internal/pkg/nodemanager/clinode"
	"github.com/tarmalonchik/speedtest/internal/pkg/nodemanager/servnode"
)

type ClientsContainer struct {
	cliNode     *clinode.ClientNodeManager
	srvNode     *servnode.ServerNodes
	measurement *measurement.Measurement
}

func getClients(_ context.Context, _ *config.Config) (*ClientsContainer, error) {
	clients := &ClientsContainer{}
	clients.cliNode = clinode.NewClientNodeManager()
	clients.srvNode = servnode.NewServerNodes()
	clients.measurement = measurement.NewMeasurement()
	return clients, nil
}
