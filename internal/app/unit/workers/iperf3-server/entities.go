package iperf3server

type Config struct {
	Iperf3Port string `envconfig:"IPERF3_PORT" default:"5201"`
}
