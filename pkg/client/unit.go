package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/tarmalonchik/speedtest/pkg/api/sdk"
)

type UnitClient struct {
	sdk.UnitServiceClient
	conn *grpc.ClientConn
}

func (u *UnitClient) CloseConnection() error {
	return u.conn.Close()
}

func NewUnitClient(serverAddr string) (*UnitClient, error) {
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &UnitClient{
		sdk.NewUnitServiceClient(conn),
		conn,
	}, nil
}
