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
	Server       webservice.Config
	Default      config.DefaultConfig
	Ping         pinger.Config
	Iperf3Server iperf3server.Config
	Iperf3Client iperf3client.Config
	Svc          svc.Config
}

func GetConfig(service string) (conf *Config, err error) {
	conf = &Config{}
	err = config.Load(service, conf)
	return conf, err
}
