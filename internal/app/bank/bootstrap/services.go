package bootstrap

import (
	"context"

	"github.com/tarmalonchik/speedtest/internal/app/bank/config"
	"github.com/tarmalonchik/speedtest/internal/app/bank/svc"
	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
)

type ServiceContainer struct {
	conf    *config.Config
	clients *ClientsContainer
	bankSvc *svc.Service
}

func GetServices(ctx context.Context, conf *config.Config) (*ServiceContainer, error) {
	var (
		err error
		sv  = &ServiceContainer{conf: conf}
	)

	if sv.clients, err = getClients(ctx, conf); err != nil {
		return nil, trace.FuncNameWithErrorMsg(err, "getting clients")
	}

	sv.bankSvc = svc.NewService(
		ctx,
		conf.Bank,
		sv.clients.cache,
	)
	return sv, nil
}
