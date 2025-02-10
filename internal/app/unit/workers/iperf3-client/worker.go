package iperf3client

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
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
	AvailableNodes(ctx context.Context, in *sdk.AvailableNodesRequest, opts ...grpc.CallOption) (*sdk.AvailableNodesResponse, error)
}

func NewWorker(conf Config, bankCli bankCli) *Worker {
	return &Worker{
		conf:    conf,
		bankCli: bankCli,
	}
}

func (t *Worker) Run(ctx context.Context) error {
	if !t.conf.IsClient {
		return nil
	}

	for {
		select {
		case <-ctx.Done():
			logrus.Infof("%s stopped successfull", trace.FuncName())
			return nil
		case <-time.NewTicker(t.conf.MeasurementPeriod).C:
			if err := t.run(ctx); err != nil {
				logrus.WithError(err).Errorf("worker")
			}
		}
	}
}

func (t *Worker) run(ctx context.Context) error {
	allNodes, err := t.bankCli.AvailableNodes(ctx, &sdk.AvailableNodesRequest{})
	if err != nil {
		return trace.FuncNameWithErrorMsg(err, "getting available nodes")
	}

	for i := range allNodes.Ip {
		if err = t.measureSingleNode(ctx, allNodes.Ip[i]); err != nil {
			logrus.WithError(err).Errorf("measuring node %s", allNodes.Ip[i])
		}
	}
	return nil
}

func (t *Worker) measureSingleNode(ctx context.Context, ip string) error {
	data, err := exec.CommandContext(ctx, "iperf3", "-c", ip, "-p", "5201", "-t5", "--json").Output()
	if err != nil {
		return trace.FuncNameWithErrorMsg(err, "executing command")
	}

	var payload IperfJsonOut

	if err = json.Unmarshal(data, &payload); err != nil {
		return trace.FuncNameWithErrorMsg(err, "unmarshal")
	}
	fmt.Println(payload.End.SumReceived.BitsPerSecond)
	fmt.Println(payload.End.SumSent.BitsPerSecond)
	return nil
}
