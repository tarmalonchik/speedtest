package bootstrap

import (
	"context"
	"fmt"

	"github.com/tarmalonchik/speedtest/internal/app/bank/config"
	"github.com/tarmalonchik/speedtest/internal/pkg/grpc"
	"github.com/tarmalonchik/speedtest/pkg/api/sdk"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func GetRouter(ctx context.Context, conf *config.Config, services *ServiceContainer, handlers *HandlerContainer) (*mux.Router, error) {
	r := mux.NewRouter()
	g := runtime.NewServeMux()

	err := sdk.RegisterBankServiceHandlerServer(ctx, g, handlers.grpcBank)
	if err != nil {
		return nil, fmt.Errorf("grpc handler init err %w", err)
	}

	grpcRouter := r.PathPrefix("").Subrouter()
	grpcRouter.Handle("/{grpc:[a-zA-Z0-9=\\-\\/]+}", g)
	return r, nil
}

func GetGRPC(_ context.Context, conf *config.Config, handlers *HandlerContainer) (*grpc.Service, error) {
	g, err := grpc.New(conf.GetGRPCAddr())
	if err != nil {
		return nil, fmt.Errorf("grpc construct err %w", err)
	}

	err = g.Init()
	if err != nil {
		return nil, fmt.Errorf("grpc init err %w", err)
	}

	g.RegisterService(handlers.grpcBank.GetServiceDesc(), handlers.grpcBank)
	return g, nil
}
