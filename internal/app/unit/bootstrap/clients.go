package bootstrap

import (
	"context"

	"github.com/tarmalonchik/speedtest/internal/app/unit/config"
	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
	"github.com/tarmalonchik/speedtest/pkg/client"
)

type ClientsContainer struct {
	bankClient *client.BankClient
}

func getClients(ctx context.Context, conf *config.Config) (clients *ClientsContainer, err error) {
	clients = &ClientsContainer{}
	if clients.bankClient, err = client.NewBankClient(conf.Ping.BankServerAddress); err != nil {
		return nil, trace.FuncNameWithErrorMsg(err, "create bank client")
	}
	return clients, nil
}
