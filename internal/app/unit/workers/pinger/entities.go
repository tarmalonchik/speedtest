package pinger

import (
	"time"
)

type Config struct {
	PingPeriod  time.Duration `envconfig:"PING_PERIOD" default:"2s"`
	MyIpAddress string        `envconfig:"MY_IP_ADDRESS" required:"true"`
	IsClient    bool          `envconfig:"IS_CLIENT" default:"false"`
}
