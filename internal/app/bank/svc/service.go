package svc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
)

type Service struct {
	ctx   context.Context
	conf  Config
	cache cache
}

type cache interface {
	Get(key string) (value []byte, ok bool)
	Add(key string, value []byte)
	GetByPrefix(prefix string) (value [][]byte)
}

func NewService(
	ctx context.Context,
	conf Config,
	cache cache,
) *Service {
	return &Service{
		ctx:   ctx,
		conf:  conf,
		cache: cache,
	}
}

func (s *Service) UpdateNodeInCache(_ context.Context, ip string) error {
	node := availableNode{
		IP:        ip,
		UpdatedAt: time.Now().UTC(),
	}
	nodeRaw, err := json.Marshal(node)
	if err != nil {
		return trace.FuncNameWithError(errors.New("json marshal"))
	}

	s.cache.Add(fmt.Sprintf(availableNodesTemp, ip), nodeRaw)
	return nil
}

func (s *Service) GetAvailableNodes(_ context.Context) (ips []string, err error) {
	respRaw := s.cache.GetByPrefix(availableNodesPrefix)
	var out = make([]availableNode, len(respRaw))

	for i := range respRaw {
		if err := json.Unmarshal(respRaw[i], &out[i]); err != nil {
			return nil, trace.FuncNameWithErrorMsg(err, "unmarshal json")
		}
	}

	for i := range out {
		if out[i].UpdatedAt.After(time.Now().UTC().Add(-s.conf.NodeIsAvailableTimeout)) {
			ips = append(ips, out[i].IP)
		}
	}
	return ips, nil
}

func (s *Service) AddNodesResults(_ context.Context, results []NodeResult) error {
	now := time.Now().UTC()
	for i := range results {
		savedRaw, ok := s.cache.Get(fmt.Sprintf(nodeSpeedTemp, results[i].IP))
		if ok {
			var saved NodeResult
			if err := json.Unmarshal(savedRaw, &saved); err != nil {
				logrus.WithError(trace.FuncNameWithError(err)).Errorf("unmarshal for ip: %s", results[i].IP)
				continue
			}
			if saved.CreatedAt.After(now.Add(-s.conf.MeasurementPeriod)) && saved.TotalSpeed() >= results[i].TotalSpeed() {
				continue
			}
		}
		dataRaw, err := json.Marshal(results[i])
		if err != nil {
			logrus.WithError(trace.FuncNameWithError(err)).Errorf("marshal for ip: %s", results[i].IP)
			continue
		}
		s.cache.Add(fmt.Sprintf(nodeSpeedTemp, results[i].IP), dataRaw)
	}
	return nil
}

func (s *Service) GetNodeSpeed(_ context.Context, ipAddress string) (inbound, outbound int64) {
	resultRaw, ok := s.cache.Get(fmt.Sprintf(nodeSpeedTemp, ipAddress))
	if !ok {
		return 0, 0
	}

	var result NodeResult
	if err := json.Unmarshal(resultRaw, &result); err != nil {
		logrus.WithError(trace.FuncNameWithError(err)).Errorf("unmarshal cache result for ip: %s", ipAddress)
		return 0, 0
	}

	if result.CreatedAt.Before(time.Now().UTC().Add(-s.conf.MeasurementPeriod * 2)) {
		return 0, 0
	}
	return result.InboundSpeed, result.OutboundSpeed
}
