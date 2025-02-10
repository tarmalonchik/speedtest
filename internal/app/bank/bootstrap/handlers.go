package bootstrap

import (
	"github.com/tarmalonchik/speedtest/internal/app/bank/handler/grpc"
)

// HandlerContainer - contains handlers
type HandlerContainer struct {
	grpcBank *grpc.BankSvc
}

// GetHandlers - handlers provider
func GetHandlers(services *ServiceContainer) *HandlerContainer {
	hCont := &HandlerContainer{}

	hCont.grpcBank = grpc.NewBankSvcHandler(services.bankSvc, *services.conf)
	return hCont
}
