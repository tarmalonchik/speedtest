package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

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
		IpAddress: s.conf.ExternalIP,
	})
	if err != nil {
		return 0, 0, trace.FuncNameWithErrorMsg(err, "getting node speed")
	}
	return resp.GetOutboundSpeed(), resp.GetInboundSpeed(), nil
}

// MeasureSpeed ...
// nolint
func (s *Service) MeasureSpeed(ctx context.Context, iperf3Server string) (inbound, outbound int64, err error) {
	var (
		data    []byte
		payload IperfJSONOut
	)

	if data, err = exec.CommandContext(
		ctx,
		"iperf3",
		"-c",
		iperf3Server,
		"-p",
		s.conf.Iperf3Port,
		fmt.Sprintf("-t%d", s.conf.Iperf3MeasurementCount),
		"--json",
	).Output(); err != nil {
		return 0, 0, trace.FuncNameWithErrorMsg(err, "measuring error")
	}

	if err = json.Unmarshal(data, &payload); err != nil {
		return 0, 0, trace.FuncNameWithErrorMsg(err, "unmarshal")
	}

	return int64(payload.End.SumReceived.BitsPerSecond), int64(payload.End.SumSent.BitsPerSecond), nil
}
