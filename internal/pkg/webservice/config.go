package webservice

import (
	"fmt"
)

type Config struct {
	HTTPPort string `envconfig:"HTTP_PORT" default:"8080"`
	GRPCPort string `envconfig:"GRPC_PORT" default:"8081"`
	Host     string `envconfig:"HOST" default:""`
}

func (c *Config) GetGRPCAddr() string { return fmt.Sprintf("%s:%s", c.Host, c.GRPCPort) }
func (c *Config) GetHTTPAddr() string { return fmt.Sprintf("%s:%s", c.Host, c.HTTPPort) }
