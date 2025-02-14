package servnode

import (
	"sync"
	"time"

	"github.com/tarmalonchik/speedtest/internal/app/bank/svc"
)

type ServerNodes struct {
	mp map[string]svc.Node
	sync.Mutex
}

func NewServerNodes() *ServerNodes {
	return &ServerNodes{
		mp: make(map[string]svc.Node),
	}
}

func (s *ServerNodes) PingNode(externalIP, internalIP, provider string) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.mp[externalIP]; ok {
		s.pingNode(externalIP, internalIP, provider)
		return
	}
	s.mp[externalIP] = svc.Node{
		ExternalIP: externalIP,
		LastUpdate: time.Now().UTC(),
		InternalIP: internalIP,
		Provider:   provider,
	}
}

func (s *ServerNodes) pingNode(externalIP, internalIP, provider string) {
	val := svc.Node{
		ExternalIP: externalIP,
		LastUpdate: time.Now().UTC(),
		InternalIP: internalIP,
		Provider:   provider,
	}
	s.mp[externalIP] = val
}

func (s *ServerNodes) GetNodes(pingPeriod time.Duration) []svc.Node {
	s.Lock()
	defer s.Unlock()

	out := make([]svc.Node, 0, len(s.mp))
	for _, val := range s.mp {
		if val.LastUpdate.UTC().Before(time.Now().UTC().Add(-3 * pingPeriod)) {
			continue
		}
		out = append(out, val)
	}
	return out
}
