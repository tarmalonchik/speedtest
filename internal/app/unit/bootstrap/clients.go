package bootstrap

import (
	"context"
	"fmt"

	"github.com/tarmalonchik/speedtest/internal/app/unit/config"
	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
	"github.com/tarmalonchik/speedtest/pkg/client"
)

type ClientsContainer struct {
	bankClient *client.BankClient
}

func getClients(ctx context.Context, conf *config.Config) (clients *ClientsContainer, err error) {
	clients = &ClientsContainer{}
	if clients.bankClient, err = client.NewBankClient(fmt.Sprintf("%s:%s", conf.Svc.BankHost, conf.Svc.BankPort)); err != nil {
		return nil, trace.FuncNameWithErrorMsg(err, "create bank client")
	}
	return clients, nil
}
