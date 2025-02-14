package clinode

import (
	"sync"
	"time"

	"github.com/tarmalonchik/speedtest/internal/app/bank/svc"
)

type node struct {
	Val  svc.Node
	Next *node
}

type ClientNodeManager struct {
	current *node
	mp      map[string]*node
	sync.Mutex
}

func NewClientNodeManager() *ClientNodeManager {
	return &ClientNodeManager{
		mp: make(map[string]*node),
	}
}

func (n *ClientNodeManager) GetClientsCount() (count int) {
	n.Lock()
	defer n.Unlock()
	return len(n.mp)
}

func (n *ClientNodeManager) PingNode(externalIP, internalIP, provider string) {
	n.Lock()
	defer n.Unlock()

	_, ok := n.mp[externalIP]
	if ok {
		n.pingNode(externalIP, internalIP, provider)
		return
	}

	if n.current == nil {
		n.current = &node{
			Val: svc.Node{
				ExternalIP: externalIP,
				InternalIP: internalIP,
				LastUpdate: time.Now().UTC(),
				Provider:   provider,
			},
		}
		n.current.Next = n.current
		n.mp[externalIP] = n.current
		return
	}
	oldNext := n.current.Next
	newNext := &node{
		Val: svc.Node{
			ExternalIP: externalIP,
			InternalIP: internalIP,
			LastUpdate: time.Now().UTC(),
			Provider:   provider,
		},
		Next: oldNext,
	}
	n.current.Next = newNext
	n.mp[externalIP] = newNext
}

func (n *ClientNodeManager) pingNode(externalIP, internalIP, provider string) {
	n.mp[externalIP].Val.LastUpdate = time.Now().UTC()
	n.mp[externalIP].Val.InternalIP = internalIP
	n.mp[externalIP].Val.Provider = provider
}

func (n *ClientNodeManager) GoNext(pingPeriod time.Duration) (svc.Node, bool) {
	n.Lock()
	defer n.Unlock()

	if n.current == nil {
		return svc.Node{}, false
	}

	out := n.current
	n.current = n.current.Next
	return out.Val, out.Val.LastUpdate.After(time.Now().UTC().Add(-pingPeriod * 3))
}
