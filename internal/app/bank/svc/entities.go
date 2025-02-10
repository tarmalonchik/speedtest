package svc

import (
	"time"
)

const (
	availableNodesPrefix = "available-node-"
	availableNodesTemp   = availableNodesPrefix + "%s"
)

type availableNode struct {
	IP        string    `json:"ip"`
	UpdatedAt time.Time `json:"updated_at"`
}
