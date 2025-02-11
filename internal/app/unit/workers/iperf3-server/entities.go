package iperf3server

type Config struct {
	IsClient   bool   `envconfig:"IS_CLIENT" default:"false"`
	Iperf3Port string `envconfig:"IPERF3_PORT" default:"5201"`
}
