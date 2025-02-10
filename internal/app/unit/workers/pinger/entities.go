package pinger

import (
	"time"
)

type Config struct {
	PingPeriod        time.Duration `envconfig:"PING_PERIOD" default:"2s"`
	BankServerAddress string        `envconfig:"BANK_SERVER_ADDRESS" default:"127.0.0.1:8081"`
	MyIpAddress       string        `envconfig:"MY_IP_ADDRESS" default:"165.23.44.12"`
	IsClient          bool          `envconfig:"IS_CLIENT" default:"false"`
}
