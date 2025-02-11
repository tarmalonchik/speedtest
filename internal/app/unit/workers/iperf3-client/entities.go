package iperf3client

import (
	"time"
)

type Config struct {
	MeasurementPeriod  time.Duration `envconfig:"MEASUREMENT_PERIOD" default:"600s"`
	BankServerAddress  string        `envconfig:"BANK_SERVER_ADDRESS" default:"127.0.0.1:8081"`
	IsClient           bool          `envconfig:"IS_CLIENT" default:"false"`
	MeasurementRetries uint8         `envconfig:"MEASUREMENT_RETRIES" default:"10"`
	Iperf3Port         string        `envconfig:"IPERF3_PORT" default:"5201"`
	MyIpAddress        string        `envconfig:"MY_IP_ADDRESS" required:"true"`
}

type speed struct {
	ipAddress    string
	inboundBits  int64
	outboundBits int64
	createdAt    time.Time
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
