package bootstrap

import (
	"github.com/tarmalonchik/speedtest/internal/app/unit/handler"
)

type HandlerContainer struct {
	handler *handler.Handler
}

func GetHandlers(services *ServiceContainer) *HandlerContainer {
	hCont := &HandlerContainer{}

	hCont.handler = handler.NewHandler(services.svc)
	return hCont
}
