package pinger

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
	"github.com/tarmalonchik/speedtest/pkg/api/sdk"
)

type Worker struct {
	conf    Config
	bankCli bankCli
}

type bankCli interface {
	Ping(ctx context.Context, in *sdk.PingRequest, opts ...grpc.CallOption) (*sdk.PingResponse, error)
}

func NewWorker(conf Config, bankCli bankCli) *Worker {
	return &Worker{
		conf:    conf,
		bankCli: bankCli,
	}
}

func (t *Worker) Run(ctx context.Context) error {
	t.run(ctx)
	for {
		select {
		case <-ctx.Done():
			logrus.Infof("%s stopped successfull", trace.FuncName())
			return nil
		case <-time.NewTicker(t.conf.PingPeriod).C:
			t.run(ctx)
		}
	}
}

func (t *Worker) run(ctx context.Context) {
	_, err := t.bankCli.Ping(ctx, &sdk.PingRequest{
		ExternalIpAddress: t.conf.ExternalIP,
		InternalIpAddress: t.conf.InternalIP,
		IsClient:          t.conf.IsClient,
	})
	if err != nil {
		logrus.WithError(trace.FuncNameWithError(err)).Errorf("sending ping request")
	} else {
		logrus.Infof("ping success")
	}
}
