package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vkidmode/server-core/pkg/core"

	"github.com/tarmalonchik/speedtest/internal/app/unit/bootstrap"
	"github.com/tarmalonchik/speedtest/internal/app/unit/config"
	"github.com/tarmalonchik/speedtest/internal/pkg/version"
	"github.com/tarmalonchik/speedtest/internal/pkg/webservice"
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
	if err = conf.ParseBase64Info(); err != nil {
		logrus.Errorf("failed to parse server data: %v", err)
		return
	}

	if conf.Ping.CurrentServerConfig.SpeedtestIsServer {
		logrus.Infof("SERVER MODE ON")
	} else {
		logrus.Infof("CLIENT MODE ON")
	}

	services, err := bootstrap.GetServices(ctx, conf)
	if err != nil {
		logrus.Errorf("failed to initiate service locator: %v", err)
		return
	}

	handlers := bootstrap.GetHandlers(services)
	router, err := bootstrap.GetRouter(handlers)
	if err != nil {
		logrus.Errorf("failed to initiate routers")
		return
	}

	ws := webservice.NewWebService(conf.Server, router)
	app := core.NewCore(nil, conf.Default.GracefulTimeout, 50)

	if conf.Ping.CurrentServerConfig.SpeedtestIsServer {
		app.AddRunner(services.GetIperf3ServerWorker().Run, true)
	} else {
		app.AddRunner(ws.Run, false)
	}
	app.AddRunner(services.GetPingWorker().Run, true)

	err = app.Launch(ctx)
	if err != nil {
		logrus.Errorf("application error: %v", err)
	}
}
