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

	conf.ParseServerModeIP()

	if conf.Iperf3Client.IsClient {
		logrus.Infof("CLIENT MODE ON")
	} else {
		logrus.Infof("SERVER MODE ON")
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
	app.AddRunner(ws.Run, false)
	if conf.Iperf3Client.IsClient {
		app.AddRunner(services.GetIperf3ClientWorker().Run, true)
	} else {
		app.AddRunner(services.GetPingWorker().Run, true)
		app.AddRunner(services.GetIperf3ServerWorker().Run, true)
	}
	err = app.Launch(ctx)
	if err != nil {
		logrus.Errorf("application error: %v", err)
	}
}
