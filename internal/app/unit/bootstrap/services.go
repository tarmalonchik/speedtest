package bootstrap

import (
	"context"

	"github.com/tarmalonchik/speedtest/internal/app/unit/config"
	"github.com/tarmalonchik/speedtest/internal/app/unit/svc"
	iperf3client "github.com/tarmalonchik/speedtest/internal/app/unit/workers/iperf3-client"
	iperf3server "github.com/tarmalonchik/speedtest/internal/app/unit/workers/iperf3-server"
	"github.com/tarmalonchik/speedtest/internal/app/unit/workers/pinger"
	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
)

type ServiceContainer struct {
	conf               *config.Config
	clients            *ClientsContainer
	pingWorker         *pinger.Worker
	iperf3ServerWorker *iperf3server.Worker
	iperf3ClientWorker *iperf3client.Worker
	svc                *svc.Service
}

func GetServices(ctx context.Context, conf *config.Config) (sv *ServiceContainer, err error) {
	sv = &ServiceContainer{conf: conf}
	if sv.clients, err = getClients(ctx, conf); err != nil {
		return nil, trace.FuncNameWithErrorMsg(err, "getting clients")
	}

	sv.pingWorker = pinger.NewWorker(conf.Ping, sv.clients.bankClient)
	sv.iperf3ServerWorker = iperf3server.NewWorker(conf.Iperf3Server)
	sv.iperf3ClientWorker = iperf3client.NewWorker(conf.Iperf3Client, sv.clients.bankClient)
	sv.svc = svc.NewService(ctx, sv.conf.Svc, sv.clients.bankClient)
	return sv, nil
}

func (s *ServiceContainer) GetPingWorker() *pinger.Worker {
	return s.pingWorker
}

func (s *ServiceContainer) GetIperf3ServerWorker() *iperf3server.Worker {
	return s.iperf3ServerWorker
}

func (s *ServiceContainer) GetIperf3ClientWorker() *iperf3client.Worker {
	return s.iperf3ClientWorker
}
