package pinger

import (
	"time"
)

type Config struct {
	PingPeriod          time.Duration `envconfig:"PING_PERIOD" default:"20s"`
	InternalIP          string        `envconfig:"INTERNAL_IP" required:"true"`
	ExternalIP          string        `envconfig:"EXTERNAL_IP" required:"true"`
	CurrentServerConfig ServerConfig
}

type ServerConfig struct {
	NodeName          string `json:"node_name"`
	LabelName         string `json:"label_name"`
	ServerCountry     string `json:"server_country"`
	ServerCapacity    string `json:"server_capacity"`
	EnableSender      string `json:"enable_sender"`
	LowPriority       string `json:"low_priority"`
	Alpha3Code        string `json:"alpha_3_code"`
	EnableIpTracker   string `json:"enable_ip_tracker"`
	IsInfrastructure  bool   `json:"is_infrastructure"`
	IpAddress         string `json:"ip_address"`
	SpeedtestIsServer bool   `json:"speedtest_is_server"`
	Provider          string `json:"provider"`
}
