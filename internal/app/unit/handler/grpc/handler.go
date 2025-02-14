package grpc

import (
	"github.com/tarmalonchik/speedtest/internal/app/unit/config"
	"github.com/tarmalonchik/speedtest/internal/app/unit/svc"
	"github.com/tarmalonchik/speedtest/internal/pkg/grpc"
	"github.com/tarmalonchik/speedtest/pkg/api/sdk"
)

type UnitSvc struct {
	grpc.Handler
	sdk.UnimplementedUnitServiceServer
	svc  *svc.Service
	conf config.Config
}

func NewUnitSvcHandler(svc *svc.Service, conf config.Config) *UnitSvc {
	return &UnitSvc{
		Handler: grpc.Handler(sdk.UnitService_ServiceDesc),
		svc:     svc,
		conf:    conf,
	}
}
