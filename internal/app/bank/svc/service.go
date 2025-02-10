package svc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

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

func (s *Service) UpdateNodeInCache(ctx context.Context, ip string) error {
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

func (s *Service) GetAvailableNodes(ctx context.Context) (ips []string, err error) {
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
