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

func (n *ClientNodeManager) AddNode(externalIP, internalIP string) {
	n.Lock()
	defer n.Unlock()

	_, ok := n.mp[externalIP]
	if ok {
		return
	}

	if n.current == nil {
		n.current = &node{
			Val: svc.Node{
				ExternalIP: externalIP,
				InternalIP: internalIP,
				LastUpdate: time.Now().UTC(),
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
		},
		Next: oldNext,
	}
	n.current.Next = newNext
	n.mp[externalIP] = newNext
}

func (n *ClientNodeManager) PingNode(externalIP, internalIP string) {
	n.Lock()
	defer n.Unlock()

	if _, ok := n.mp[externalIP]; !ok {
		return
	}
	n.mp[externalIP].Val.LastUpdate = time.Now().UTC()
	n.mp[externalIP].Val.InternalIP = internalIP
}

func (n *ClientNodeManager) GoNext() svc.Node {
	n.Lock()
	defer n.Unlock()

	out := n.current
	n.current = n.current.Next
	return out.Val
}
