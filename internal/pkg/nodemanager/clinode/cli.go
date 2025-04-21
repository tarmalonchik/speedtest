package clinode

import (
	"sync"
	"time"

	"github.com/tarmalonchik/speedtest/internal/app/bank/svc"
)

type ClientNodeManager struct {
	nodes []svc.Node
	sync.Mutex
}

func NewClientNodeManager() *ClientNodeManager {
	return &ClientNodeManager{
		nodes: make([]svc.Node, 0),
	}
}

func (n *ClientNodeManager) GetClientsCount() (count int) {
	return len(n.nodes)
}

func (n *ClientNodeManager) PingNode(externalIP, internalIP, provider string, nowTime time.Time) {
	n.Lock()
	defer n.Unlock()

	for i := range n.nodes {
		if n.nodes[i].ExternalIP == externalIP {
			n.nodes[i].InternalIP = internalIP
			n.nodes[i].Provider = provider
			n.nodes[i].LastUpdate = nowTime
			return
		}
	}
	n.nodes = append(n.nodes, svc.Node{
		ExternalIP: externalIP,
		InternalIP: internalIP,
		Provider:   provider,
		LastUpdate: nowTime,
	})
}

func (n *ClientNodeManager) GetNodes(pingPeriod time.Duration) (out []svc.Node) {
	n.Lock()
	defer n.Unlock()

	for i := range n.nodes {
		if n.nodes[i].LastUpdate.Before(time.Now().UTC().Add(-pingPeriod * 3)) {
			continue
		}
		out = append(out, n.nodes[i])
	}
	return out
}
