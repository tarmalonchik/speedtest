package bootstrap

import (
	"github.com/tarmalonchik/speedtest/internal/app/unit/handler"
	"github.com/tarmalonchik/speedtest/internal/app/unit/handler/grpc"
)

type HandlerContainer struct {
	handler     *handler.Handler
	grpcHandler *grpc.UnitSvc
}

func GetHandlers(services *ServiceContainer) *HandlerContainer {
	hCont := &HandlerContainer{}

	hCont.handler = handler.NewHandler(services.svc)
	hCont.grpcHandler = grpc.NewUnitSvcHandler(services.svc, *services.conf)
	return hCont
}
