package main

import (
	"context"

	"github.com/tarmalonchik/speedtest/internal/app/bank/bootstrap"
	"github.com/tarmalonchik/speedtest/internal/app/bank/config"
	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
	"github.com/tarmalonchik/speedtest/internal/pkg/version"
	"github.com/tarmalonchik/speedtest/internal/pkg/webservice"

	"github.com/sirupsen/logrus"
	"github.com/vkidmode/server-core/pkg/core"
)

func init() {
	if version.Service == "" {
		version.Service = "bank"
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

	handlers := bootstrap.GetHandlers(services)
	router, err := bootstrap.GetRouter(ctx, conf, services, handlers)
	if err != nil {
		logrus.Errorf("failed to initiate routers")
		return
	}

	ws := webservice.NewWebService(conf.Server, router)
	grpc, err := bootstrap.GetGRPC(ctx, conf, handlers)
	if err != nil {
		logrus.WithError(trace.FuncNameWithError(err)).Error("failed to initiate grpc")
		return
	}

	app := core.NewCore(nil, conf.Default.GracefulTimeout, 50)
	app.AddRunner(ws.Run, true)
	app.AddRunner(grpc.Run, true)

	err = app.Launch(ctx)
	if err != nil {
		logrus.Errorf("application error: %v", err)
	}
}
