package grpc

import (
	"github.com/tarmalonchik/speedtest/internal/app/bank/config"
	"github.com/tarmalonchik/speedtest/internal/app/bank/svc"
	"github.com/tarmalonchik/speedtest/internal/pkg/grpc"
	"github.com/tarmalonchik/speedtest/pkg/api/sdk"
)

type BankSvc struct {
	grpc.Handler
	sdk.UnimplementedBankServiceServer
	svc  *svc.Service
	conf config.Config
}

func NewBankSvcHandler(svc *svc.Service, conf config.Config) *BankSvc {
	return &BankSvc{
		Handler: grpc.Handler(sdk.BankService_ServiceDesc),
		svc:     svc,
		conf:    conf,
	}
}
