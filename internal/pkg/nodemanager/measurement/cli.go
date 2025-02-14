package measurement

import (
	"sync"
	"time"
)

type speed struct {
	InboundSpeed  int64
	OutboundSpeed int64
	CreatedAt     time.Time
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
		CreatedAt:     time.Now().UTC(),
	}
}

func (m *Measurement) GetData(externalIP string, period time.Duration) (inbound, outbound int64) {
	m.Lock()
	defer m.Unlock()
	resp := m.mp[externalIP]
	if resp.CreatedAt.UTC().Before(time.Now().UTC().Add(-3 * period)) {
		return 0, 0
	}
	return resp.InboundSpeed, resp.OutboundSpeed
}
