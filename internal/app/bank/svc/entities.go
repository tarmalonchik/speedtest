package svc

import (
	"time"
)

type Config struct {
	NodeIsAvailableTimeout time.Duration `envconfig:"NODE_IS_AVAILABLE_TIMEOUT" default:"40s"`
	MeasurementPeriod      time.Duration `envconfig:"MEASUREMENT_PERIOD" default:"600s"`
	UnitGRPCPort           string        `envconfig:"UNIT_GRPC_PORT" required:"true"`
	MeasurementStepsPeriod time.Duration `envconfig:"MEASUREMENTS_STEPS_PERIOD" required:"true"`
}

type Node struct {
	InternalIP string
	ExternalIP string
	LastUpdate time.Time
	Provider   string
}

type speed struct {
	InboundSpeed     int64
	OutboundSpeed    int64
	ServerExternalIP string
}
