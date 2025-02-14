package svc

type Config struct {
	BankHost               string `envconfig:"BANK_HOST" required:"true"`
	BankPort               string `envconfig:"BANK_PORT" required:"true"`
	InternalIP             string `envconfig:"INTERNAL_IP" required:"true"`
	ExternalIP             string `envconfig:"EXTERNAL_IP" required:"true"`
	Iperf3Port             string `envconfig:"IPERF3_PORT" default:"5201"`
	Iperf3MeasurementCount int    `envconfig:"IPERF3_MEASUREMENT_COUNT" required:"true"`
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
