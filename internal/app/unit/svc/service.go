package svc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
	"github.com/tarmalonchik/speedtest/pkg/api/sdk"
)

type bankCli interface {
	GetNodeSpeed(ctx context.Context, in *sdk.GetNodeSpeedRequest, opts ...grpc.CallOption) (*sdk.GetNodeSpeedResponse, error)
}

type Service struct {
	ctx     context.Context
	conf    Config
	bankCli bankCli
}

func NewService(
	ctx context.Context,
	conf Config,
	bankCli bankCli,
) *Service {
	return &Service{
		ctx:     ctx,
		conf:    conf,
		bankCli: bankCli,
	}
}

func (s *Service) GetNodeSpeed(ctx context.Context) (outbound, inbound int64, err error) {
	resp, err := s.bankCli.GetNodeSpeed(ctx, &sdk.GetNodeSpeedRequest{
		IpAddress: s.conf.MyIpAddress,
	})
	if err != nil {
		return 0, 0, trace.FuncNameWithErrorMsg(err, "getting node speed")
	}
	return resp.OutboundSpeed, resp.InboundSpeed, nil
}
