package iperf3client

import (
	"context"
	"encoding/json"
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
	AddNodesResults(ctx context.Context, in *sdk.AddNodesResultsRequest, opts ...grpc.CallOption) (*sdk.AddNodesResultsResponse, error)
}

func NewWorker(conf Config, bankCli bankCli) *Worker {
	return &Worker{
		conf:    conf,
		bankCli: bankCli,
	}
}

func (t *Worker) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			logrus.Infof("%s stopped successfull", trace.FuncName())
			return nil
		case <-time.NewTicker(t.conf.MeasurementPeriod).C:
			if err := t.run(ctx); err != nil {
				logrus.WithError(trace.FuncNameWithError(err)).Errorf("worker")
			}
		}
	}
}

func (t *Worker) run(ctx context.Context) error {
	allNodes, err := t.bankCli.AvailableNodes(ctx, &sdk.AvailableNodesRequest{})
	if err != nil {
		return trace.FuncNameWithErrorMsg(err, "getting available nodes")
	}

	var allSpeeds []speed

	for i := range allNodes.Ip {
		singleSpeed, err := t.measureSingleNode(ctx, allNodes.Ip[i])
		if err != nil {
			logrus.WithError(trace.FuncNameWithError(err)).Errorf("measuring node %s", allNodes.Ip[i])
		} else {
			logrus.Infof("measured success ip %s inbound %d outbound %d", allNodes.Ip[i], singleSpeed.inboundBits, singleSpeed.outboundBits)
			allSpeeds = append(allSpeeds, singleSpeed)
		}
	}

	currentNodeSpeed := t.getMaxSpeed(allSpeeds)
	currentNodeSpeed.ipAddress = t.conf.MyIpAddress
	currentNodeSpeed.createdAt = time.Now().UTC()
	allSpeeds = append(allSpeeds, currentNodeSpeed)

	bankCliRequest := &sdk.AddNodesResultsRequest{
		Items: make([]*sdk.AddNodesResultsRequestItem, len(allSpeeds)),
	}

	for i := range allSpeeds {
		bankCliRequest.Items[i] = &sdk.AddNodesResultsRequestItem{
			IpAddress:     allSpeeds[i].ipAddress,
			InboundSpeed:  allSpeeds[i].inboundBits,
			OutboundSpeed: allSpeeds[i].outboundBits,
			CreatedAt:     allSpeeds[i].createdAt.Unix(),
		}
	}

	if _, err = t.bankCli.AddNodesResults(ctx, bankCliRequest); err != nil {
		return trace.FuncNameWithErrorMsg(err, "sending results")
	}
	return nil
}

func (t *Worker) getMaxSpeed(allSpeeds []speed) speed {
	if len(allSpeeds) == 0 {
		return speed{}
	}
	index := 0
	maxSpeed := int64(0)

	for i := range allSpeeds {
		sp := allSpeeds[i].outboundBits + allSpeeds[i].inboundBits
		if sp > maxSpeed {
			maxSpeed = sp
			index = i
		}
	}
	return allSpeeds[index]
}

func (t *Worker) measureSingleNode(ctx context.Context, ip string) (out speed, err error) {
	var (
		data    []byte
		payload IperfJsonOut
	)

	for i := 0; i < int(t.conf.MeasurementRetries)+1; i++ {
		data, err = exec.CommandContext(ctx, "iperf3", "-c", ip, "-p", t.conf.Iperf3Port, "-t5", "--json").Output()
		if err != nil {
			logrus.WithError(trace.FuncNameWithError(err)).Errorf("measuring node %s", ip)
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	if err = json.Unmarshal(data, &payload); err != nil {
		return speed{}, trace.FuncNameWithErrorMsg(err, "unmarshal")
	}

	return speed{
		ipAddress:    ip,
		inboundBits:  int64(payload.End.SumReceived.BitsPerSecond),
		outboundBits: int64(payload.End.SumSent.BitsPerSecond),
		createdAt:    time.Now().UTC(),
	}, nil
}
