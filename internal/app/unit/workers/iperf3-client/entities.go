package iperf3client

import (
	"time"
)

type Config struct {
	MeasurementPeriod time.Duration `envconfig:"MEASUREMENT_PERIOD" default:"2s"`
	BankServerAddress string        `envconfig:"BANK_SERVER_ADDRESS" default:"127.0.0.1:8081"`
	IsClient          bool          `envconfig:"IS_CLIENT" default:"false"`
}

type speed struct {
	inboundBits  int64
	outboundBits int64
}

type IperfJsonOut struct {
	End End `json:"end"`
}

type End struct {
	SumSent     Send     `json:"sum_sent"`
	SumReceived Received `json:"sum_received"`
}

type Received struct {
	BitsPerSecond float64 `json:"bits_per_second"`
}

type Send struct {
	BitsPerSecond float64 `json:"bits_per_second"`
}
