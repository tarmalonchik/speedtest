package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/tarmalonchik/speedtest/pkg/api/sdk"
)

type BankClient struct {
	sdk.BankServiceClient
}

func NewBankClient(serverAddr string) (*BankClient, error) {
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &BankClient{sdk.NewBankServiceClient(conn)}, nil
}
