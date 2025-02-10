package iperf3client

import (
	"time"
)

type Config struct {
	MeasurementPeriod time.Duration `envconfig:"MEASUREMENT_PERIOD" default:"2s"`
	BankServerAddress string        `envconfig:"BANK_SERVER_ADDRESS" default:"127.0.0.1:8081"`
	IsClient          bool          `envconfig:"IS_CLIENT" default:"false"`
}
