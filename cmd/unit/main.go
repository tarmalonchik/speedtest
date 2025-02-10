package main

import (
	"context"

	"github.com/tarmalonchik/speedtest/internal/app/unit/bootstrap"
	"github.com/tarmalonchik/speedtest/internal/app/unit/config"
	"github.com/tarmalonchik/speedtest/internal/pkg/version"

	"github.com/sirupsen/logrus"
	"github.com/vkidmode/server-core/pkg/core"
)

func init() {
	if version.Service == "" {
		version.Service = "unit"
	}
}

func main() {
	ctx := context.Background()

	conf, err := config.GetConfig(version.Service)
	if err != nil {
		logrus.Errorf("failed to load environment: %v", err)
		return
	}

	services, err := bootstrap.GetServices(ctx, conf)
	if err != nil {
		logrus.Errorf("failed to initiate service locator: %v", err)
		return
	}

	app := core.NewCore(nil, conf.GracefulTimeout, 50)
	app.AddRunner(services.GetPingWorker().Run, true)
	app.AddRunner(services.GetIperf3ServerWorker().Run, true)
	app.AddRunner(services.GetIperf3ClientWorker().Run, true)
	err = app.Launch(ctx)
	if err != nil {
		logrus.Errorf("application error: %v", err)
	}
}
