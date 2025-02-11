package config

import (
	"github.com/tarmalonchik/speedtest/internal/app/unit/svc"
	iperf3client "github.com/tarmalonchik/speedtest/internal/app/unit/workers/iperf3-client"
	iperf3server "github.com/tarmalonchik/speedtest/internal/app/unit/workers/iperf3-server"
	"github.com/tarmalonchik/speedtest/internal/app/unit/workers/pinger"
	"github.com/tarmalonchik/speedtest/internal/pkg/config"
	"github.com/tarmalonchik/speedtest/internal/pkg/webservice"
)

// Config contains all environment variables
type Config struct {
	Server             webservice.Config
	Default            config.DefaultConfig
	Ping               pinger.Config
	Iperf3Server       iperf3server.Config
	Iperf3Client       iperf3client.Config
	Svc                svc.Config
	EnableClientModeIP []string `envconfig:"CLIENT_MODE_LIST_JSON" required:"true"`
}

func (c *Config) ParseServerModeIP() {
	for i := range c.EnableClientModeIP {
		if c.EnableClientModeIP[i] == c.Iperf3Client.MyIpAddress {
			c.Iperf3Server.IsClient = true
			c.Iperf3Client.IsClient = true
			c.Ping.IsClient = true
			return
		}
	}
	c.Iperf3Server.IsClient = false
	c.Iperf3Client.IsClient = false
	c.Ping.IsClient = false
}

func GetConfig(service string) (conf *Config, err error) {
	conf = &Config{}
	err = config.Load(service, conf)
	return conf, err
}
