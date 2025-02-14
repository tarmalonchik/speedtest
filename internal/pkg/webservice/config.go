package webservice

import (
	"fmt"
)

type Config struct {
	HTTPPort string `envconfig:"HTTP_PORT" required:"true"`
	GRPCPort string `envconfig:"GRPC_PORT" required:"true"`
	Host     string `envconfig:"HOST" default:""`
}

func (c *Config) GetGRPCAddr() string { return fmt.Sprintf("%s:%s", c.Host, c.GRPCPort) }
func (c *Config) GetHTTPAddr() string { return fmt.Sprintf("%s:%s", c.Host, c.HTTPPort) }
