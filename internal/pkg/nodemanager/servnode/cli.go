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

func (s *ServerNodes) AddNode(externalIP, internalIP string) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.mp[externalIP]; ok {
		return
	}
	s.mp[externalIP] = svc.Node{
		ExternalIP: externalIP,
		LastUpdate: time.Now().UTC(),
		InternalIP: internalIP,
	}
}

func (s *ServerNodes) PingNode(externalIP, internalIP string) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.mp[externalIP]; !ok {
		return
	}

	val := svc.Node{
		ExternalIP: externalIP,
		LastUpdate: time.Now().UTC(),
		InternalIP: internalIP,
	}
	s.mp[externalIP] = val
}

func (s *ServerNodes) GetNodes() []svc.Node {
	s.Lock()
	defer s.Unlock()

	out := make([]svc.Node, 0, len(s.mp))
	for _, val := range s.mp {
		//if val.LastUpdate.Before() // todo skip old ones based on ping time
		out = append(out, val)
	}
	return out
}
