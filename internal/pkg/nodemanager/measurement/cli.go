package measurement

import (
	"sync"
)

type speed struct {
	InboundSpeed  int64
	OutboundSpeed int64
}

type Measurement struct {
	sync.Mutex
	mp map[string]speed
}

func NewMeasurement() *Measurement {
	return &Measurement{
		mp: make(map[string]speed),
	}
}

func (m *Measurement) AddData(externalIP string, inbound, outbound int64) {
	m.Lock()
	defer m.Unlock()
	m.mp[externalIP] = speed{
		InboundSpeed:  inbound,
		OutboundSpeed: outbound,
	}
}

func (m *Measurement) GetData(externalIP string) (inbound, outbound int64) {
	m.Lock()
	defer m.Unlock()
	resp := m.mp[externalIP]
	return resp.InboundSpeed, resp.OutboundSpeed
}
