package bootstrap

import (
	"context"

	"github.com/tarmalonchik/speedtest/internal/app/bank/config"
	"github.com/tarmalonchik/speedtest/internal/app/bank/svc"
)

type ServiceContainer struct {
	conf    *config.Config
	clients *ClientsContainer
	bankSvc *svc.Service
}

func GetServices(ctx context.Context, conf *config.Config) *ServiceContainer {
	var (
		sv = &ServiceContainer{conf: conf}
	)

	sv.clients = getClients(ctx, conf)

	sv.bankSvc = svc.NewService(
		ctx,
		conf.Bank,
		sv.clients.srvNode,
		sv.clients.cliNode,
		sv.clients.measurement,
	)
	return sv
}

func (s *ServiceContainer) GetMeasurementWorker() *svc.Service {
	return s.bankSvc
}
