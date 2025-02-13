package pinger

import (
	"time"
)

type Config struct {
	PingPeriod time.Duration `envconfig:"PING_PERIOD" default:"20s"`
	InternalIP string        `envconfig:"INTERNAL_IP" required:"true"`
	ExternalIP string        `envconfig:"EXTERNAL_IP" required:"true"`
	IsClient   bool          `envconfig:"IS_CLIENT" default:"false"`
}
