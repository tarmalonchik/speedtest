package config

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/tarmalonchik/speedtest/internal/app/unit/svc"
	iperf3server "github.com/tarmalonchik/speedtest/internal/app/unit/workers/iperf3-server"
	"github.com/tarmalonchik/speedtest/internal/app/unit/workers/pinger"
	"github.com/tarmalonchik/speedtest/internal/pkg/config"
	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
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
	CurrentServerConfig  pinger.ServerConfig
}

func (c *Config) ParseBase64Info() error {
	rawData, err := base64.StdEncoding.DecodeString(c.Base64AllServersData)
	if err != nil {
		return trace.FuncNameWithErrorMsg(err, "error while parsing")
	}
	stringData := string(rawData)
	stringData = strings.ReplaceAll(stringData, "'", "\"")
	stringData = strings.ReplaceAll(stringData, "False", "false")
	stringData = strings.ReplaceAll(stringData, "True", "true")

	var servers []pinger.ServerConfig

	if err = json.Unmarshal([]byte(stringData), &servers); err != nil {
		return trace.FuncNameWithErrorMsg(err, "unmarshal")
	}

	for i := range servers {
		if c.Ping.ExternalIP == servers[i].IPAddress {
			c.CurrentServerConfig = servers[i]
			c.Ping.CurrentServerConfig = servers[i]
			return nil
		}
	}
	return trace.FuncNameWithError(errors.New("server not found"))
}

func GetConfig(service string) (conf *Config, err error) {
	conf = &Config{}
	err = config.Load(service, conf)
	return conf, err
}
