package iperf3server

type Config struct {
	IsClient bool `envconfig:"IS_CLIENT" default:"false"`
}
