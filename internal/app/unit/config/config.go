package config

import (
	"github.com/tarmalonchik/speedtest/internal/app/unit/svc"
	iperf3server "github.com/tarmalonchik/speedtest/internal/app/unit/workers/iperf3-server"
	"github.com/tarmalonchik/speedtest/internal/app/unit/workers/pinger"
	"github.com/tarmalonchik/speedtest/internal/pkg/config"
	"github.com/tarmalonchik/speedtest/internal/pkg/webservice"
)

// Config contains all environment variables
type Config struct {
	Server               webservice.Config
	Default              config.DefaultConfig
	Ping                 pinger.Config
	Iperf3Server         iperf3server.Config
	Svc                  svc.Config
	Base64AllServersData string `envconfig:"BASE64_ALL_SERVERS_data" required:"true"`
}

func (c *Config) ParseServerModeIP() {

	for i := range c.EnableServerModeIP {
		if c.EnableServerModeIP[i] == c.Svc.ExternalIP {
			c.Ping.IsClient = false
			return
		}
	}
	c.Ping.IsClient = true
}

func GetConfig(service string) (conf *Config, err error) {
	conf = &Config{}
	err = config.Load(service, conf)
	return conf, err
}
