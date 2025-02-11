package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"

	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
)

type DefaultConfig struct {
	GracefulTimeout time.Duration `envconfig:"GRACEFUL_TIMEOUT" default:"40s"`
}

func Load(service string, conf interface{}) error {
	configFileNames := make([]string, 0, 1)

	path := fmt.Sprintf("./configs/%s/env", service)
	if _, err := os.Stat(path); err == nil {
		configFileNames = append(configFileNames, path)
	}

	if err := loadConfig(conf, "", configFileNames...); err != nil {
		return trace.FuncNameWithErrorMsg(err, "load environment")
	}
	return nil
}

func loadConfig(c interface{}, prefix string, filenames ...string) error {
	err := godotenv.Load(filenames...)
	if err != nil {
		log.WithField("filenames", filenames).Info("config file not found, using defaults")
	}

	err = envconfig.Process(prefix, c)
	if err != nil {
		return fmt.Errorf("error env config loading: %w", err)
	}

	return nil
}
