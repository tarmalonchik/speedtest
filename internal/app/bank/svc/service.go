package svc

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
	"github.com/tarmalonchik/speedtest/pkg/api/sdk"
	"github.com/tarmalonchik/speedtest/pkg/client"
)

type cliNodeManager interface {
	AddNode(externalIP, internalIP string)
	PingNode(externalIP, internalIP, provider string)
	GoNext() Node
}

type serverNodeManager interface {
	AddNode(externalIP, internalIP string)
	PingNode(externalIP, internalIP, provider string)
	GetNodes() []Node
}

type measurementManager interface {
	AddData(externalIP string, inbound, outbound int64)
	GetData(externalIP string) (inbound, outbound int64)
}

type Service struct {
	ctx                context.Context
	conf               Config
	clientNodeManager  cliNodeManager
	serverNodeManager  serverNodeManager
	measurementManager measurementManager
}

func NewService(
	ctx context.Context,
	conf Config,
	sNode serverNodeManager,
	cNode cliNodeManager,
	measurementManager measurementManager,
) *Service {
	svc := &Service{
		ctx:                ctx,
		conf:               conf,
		clientNodeManager:  cNode,
		serverNodeManager:  sNode,
		measurementManager: measurementManager,
	}
	return svc
}

func (s *Service) Ping(_ context.Context, externalIP, internalIP string, isClient bool, provider string) {
	if isClient {
		s.clientNodeManager.PingNode(externalIP, internalIP, provider)
		return
	}
	s.serverNodeManager.PingNode(externalIP, internalIP, provider)
}

func (s *Service) AddNode(_ context.Context, externalIP, internalIP string, isClient bool) {
	if isClient {
		s.clientNodeManager.AddNode(externalIP, internalIP)
		return
	}
	s.serverNodeManager.AddNode(externalIP, internalIP)
}

func (s *Service) GetNodeSpeed(_ context.Context, externalIP string) (inbound, outbound int64) {
	return s.measurementManager.GetData(externalIP)
}

func (s *Service) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			logrus.Infof("%s stopped successfull", trace.FuncName())
			return nil
		case <-time.NewTicker(s.conf.MeasurementStepsPeriod).C:
			if err := s.run(ctx); err != nil {
				logrus.WithError(trace.FuncNameWithError(err)).Errorf("runnng")
			}
		}
	}
}

func (s *Service) run(ctx context.Context) error {
	unit := s.clientNodeManager.GoNext()

	unitClient, err := client.NewUnitClient(fmt.Sprintf("%s:%s", unit.InternalIP, s.conf.UnitGRPCPort))
	if err != nil {
		return trace.FuncNameWithErrorMsg(err, "create unit client")
	}
	defer func() {
		_ = unitClient.CloseConnection()
	}()

	serverNodes := s.serverNodeManager.GetNodes()

	var speeds = make([]speed, 0, len(serverNodes))

	for i := range serverNodes {
		measureResp, err := unitClient.Measure(ctx, &sdk.MeasureRequest{
			Iperf3ServerIp: serverNodes[i].ExternalIP,
		})
		if err != nil {
			logrus.WithError(trace.FuncNameWithError(err)).Errorf("measuring node")
			continue
		}

		speeds = append(speeds, speed{
			InboundSpeed:     measureResp.InboundSpeed,
			OutboundSpeed:    measureResp.OutboundSpeed,
			ServerExternalIP: serverNodes[i].ExternalIP,
		})
	}

	if len(speeds) == 0 {
		return nil
	}

	maxSpeed := getMaxSpeed(speeds)
	s.measurementManager.AddData(unit.ExternalIP, maxSpeed.InboundSpeed, maxSpeed.OutboundSpeed)
	s.measurementManager.AddData(maxSpeed.ServerExternalIP, maxSpeed.InboundSpeed, maxSpeed.OutboundSpeed)
	return nil
}

func getMaxSpeed(allSpeeds []speed) speed {
	index := 0
	maxSpeed := int64(0)

	for i := range allSpeeds {
		sp := allSpeeds[i].OutboundSpeed + allSpeeds[i].InboundSpeed
		if sp > maxSpeed {
			maxSpeed = sp
			index = i
		}
	}
	return allSpeeds[index]
}
