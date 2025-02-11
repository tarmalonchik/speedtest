package svc

import (
	"time"
)

const (
	availableNodesPrefix = "available-node-"
	availableNodesTemp   = availableNodesPrefix + "%s"
	nodeSpeedTemp        = "node-speed-%s"
)

type Config struct {
	NodeIsAvailableTimeout time.Duration `envconfig:"NODE_IS_AVAILABLE_TIMEOUT" default:"40s"`
	MeasurementPeriod      time.Duration `envconfig:"MEASUREMENT_PERIOD" default:"600s"`
}

type availableNode struct {
	IP        string    `json:"ip"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NodeResult struct {
	IP            string    `json:"ip"`
	InboundSpeed  int64     `json:"inbound_speed"`
	OutboundSpeed int64     `json:"outbound_speed"`
	CreatedAt     time.Time `json:"created_at"`
}

func (n *NodeResult) TotalSpeed() int64 {
	return n.InboundSpeed + n.OutboundSpeed
}
