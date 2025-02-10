package config

import (
	"github.com/tarmalonchik/speedtest/internal/app/bank/svc"
	"github.com/tarmalonchik/speedtest/internal/pkg/config"
)

// Config contains all environment variables
type Config struct {
	config.DefaultConfig
	Bank svc.Config
}

func GetConfig(service string) (conf *Config, err error) {
	conf = &Config{}
	err = config.Load(service, conf)
	return conf, err
}
