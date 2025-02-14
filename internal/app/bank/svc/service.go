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
	PingNode(externalIP, internalIP, provider string, nowTime time.Time)
	GetNodes(pingPeriod time.Duration) (out []Node)
	GetClientsCount() (count int)
}

type serverNodeManager interface {
	PingNode(externalIP, internalIP, provider string)
	GetNodes(pingPeriod time.Duration) []Node
}

type measurementManager interface {
	AddData(externalIP string, inbound, outbound int64)
	GetData(externalIP string, measurementPeriod time.Duration) (inbound, outbound int64)
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
		s.clientNodeManager.PingNode(externalIP, internalIP, provider, time.Now().UTC())
		return
	}
	s.serverNodeManager.PingNode(externalIP, internalIP, provider)
}

func (s *Service) GetNodeSpeed(_ context.Context, externalIP string) (inbound, outbound int64) {
	return s.measurementManager.GetData(externalIP, s.conf.MeasurementPeriod)
}

func (s *Service) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			logrus.Infof("%s stopped successfull", trace.FuncName())
			return nil
		case <-time.NewTicker(s.conf.MeasurementPeriod).C:
			if err := s.run(ctx); err != nil {
				logrus.WithError(trace.FuncNameWithError(err)).Errorf("runnng")
			}
		}
	}
}

func (s *Service) run(ctx context.Context) error {
	units := s.clientNodeManager.GetNodes(s.conf.PingPeriod)
	servers := s.serverNodeManager.GetNodes(s.conf.PingPeriod)

	mpServersSpeed := make(map[string]speed)
	mpUnitsSpeed := make(map[string]speed)

	for i := range units {
		for j := range servers {
			if s.conf.EnableInProviderBan {
				if units[i].Provider == servers[j].Provider {
					continue
				}
			}

			measuredSpeed, err := s.measureSingleNode(ctx, units[i].InternalIP, servers[j].ExternalIP)
			if err != nil {
				logrus.WithError(trace.FuncNameWithError(err)).Errorf("measuring from client:%s to server:%s ",
					units[i].ExternalIP, servers[j].ExternalIP)
				continue
			} else {
				logrus.Infof("success measuring from client:%s to server:%s inbound: %d outbound: %d",
					units[i].ExternalIP, servers[j].ExternalIP, measuredSpeed.InboundSpeed, measuredSpeed.OutboundSpeed)
				if val, ok := mpServersSpeed[servers[j].ExternalIP]; ok {
					if val.GetSum() < measuredSpeed.GetSum() {
						mpServersSpeed[servers[j].ExternalIP] = speed{
							InboundSpeed:  measuredSpeed.InboundSpeed,
							OutboundSpeed: measuredSpeed.OutboundSpeed,
						}
					}
				} else {
					mpServersSpeed[servers[j].ExternalIP] = speed{
						InboundSpeed:  measuredSpeed.InboundSpeed,
						OutboundSpeed: measuredSpeed.OutboundSpeed,
					}
				}

				if val, ok := mpUnitsSpeed[units[i].ExternalIP]; ok {
					if val.GetSum() < measuredSpeed.GetSum() {
						mpUnitsSpeed[units[i].ExternalIP] = speed{
							InboundSpeed:  measuredSpeed.InboundSpeed,
							OutboundSpeed: measuredSpeed.OutboundSpeed,
						}
					}
				} else {
					mpUnitsSpeed[units[i].ExternalIP] = speed{
						InboundSpeed:  measuredSpeed.InboundSpeed,
						OutboundSpeed: measuredSpeed.OutboundSpeed,
					}
				}
			}
		}
	}

	for key, val := range mpServersSpeed {
		s.measurementManager.AddData(key, val.InboundSpeed, val.OutboundSpeed)
	}
	for key, val := range mpUnitsSpeed {
		s.measurementManager.AddData(key, val.InboundSpeed, val.OutboundSpeed)
	}
	return nil
}

func (s *Service) measureSingleNode(ctx context.Context, fromNodeIP, toNodeExternalIP string) (speed, error) {
	unitClient, err := client.NewUnitClient(fmt.Sprintf("%s:%s", fromNodeIP, s.conf.UnitGRPCPort))
	if err != nil {
		return speed{}, trace.FuncNameWithErrorMsg(err, "create unit client")
	}
	defer func() {
		_ = unitClient.CloseConnection()
	}()

	var measureResp = &sdk.MeasureResponse{}
	for i := 0; i < 3; i++ {
		measureResp, err = unitClient.Measure(ctx, &sdk.MeasureRequest{
			Iperf3ServerIp: toNodeExternalIP,
		})
		if err == nil {
			break
		}
	}
	if err != nil {
		return speed{}, trace.FuncNameWithErrorMsg(err, "measuring error")
	}
	if measureResp == nil {
		return speed{}, trace.FuncNameWithErrorMsg(err, "invalid response")
	}

	return speed{
		InboundSpeed:  measureResp.InboundSpeed,
		OutboundSpeed: measureResp.OutboundSpeed,
	}, nil
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
