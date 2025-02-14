package svc

import (
	"time"
)

type Config struct {
	NodeIsAvailableTimeout time.Duration `envconfig:"NODE_IS_AVAILABLE_TIMEOUT" default:"40s"`
	UnitGRPCPort           string        `envconfig:"UNIT_GRPC_PORT" required:"true"`
	MeasurementPeriod      time.Duration `envconfig:"MEASUREMENT_PERIOD" required:"true"`
	EnableInProviderBan    bool          `envconfig:"ENABLE_IN_PROVIDER_BAN" required:"true"`
	PingPeriod             time.Duration `envconfig:"PING_PERIOD" default:"20s"`
	Iperf3MeasurementCount int           `envconfig:"IPERF3_MEASUREMENT_COUNT" required:"true"`
}

type Node struct {
	InternalIP string
	ExternalIP string
	LastUpdate time.Time
	Provider   string
}

type speed struct {
	InboundSpeed  int64
	OutboundSpeed int64
}

func (s speed) GetSum() int64 {
	return s.InboundSpeed + s.OutboundSpeed
}
