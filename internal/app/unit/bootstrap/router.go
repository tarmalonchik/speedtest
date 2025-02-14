package bootstrap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/tarmalonchik/speedtest/internal/app/unit/config"
	"github.com/tarmalonchik/speedtest/internal/pkg/grpc"
	"github.com/tarmalonchik/speedtest/pkg/api/sdk"
)

func GetRouter(ctx context.Context, handlers *HandlerContainer) (*mux.Router, error) {
	r := mux.NewRouter()
	g := runtime.NewServeMux()

	err := sdk.RegisterUnitServiceHandlerServer(ctx, g, handlers.grpcHandler)
	if err != nil {
		return nil, fmt.Errorf("grpc handler init err %w", err)
	}

	r.HandleFunc("/speedtest", handlers.handler.Speedtest).Methods(http.MethodGet)
	grpcRouter := r.PathPrefix("").Subrouter()
	grpcRouter.Handle("/{grpc:[a-zA-Z0-9=\\-\\/]+}", g)
	return r, nil
}

func GetGRPC(_ context.Context, conf *config.Config, handlers *HandlerContainer) (*grpc.Service, error) {
	g, err := grpc.New(conf.Server.GetGRPCAddr())
	if err != nil {
		return nil, fmt.Errorf("grpc construct err %w", err)
	}

	err = g.Init()
	if err != nil {
		return nil, fmt.Errorf("grpc init err %w", err)
	}

	g.RegisterService(handlers.grpcHandler.GetServiceDesc(), handlers.grpcHandler)
	return g, nil
}
