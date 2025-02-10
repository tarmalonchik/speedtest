package svc

import (
	"time"
)

type Config struct {
	NodeIsAvailableTimeout time.Duration `envconfig:"NODE_IS_AVAILABLE_TIMEOUT" default:"40s"`
}
