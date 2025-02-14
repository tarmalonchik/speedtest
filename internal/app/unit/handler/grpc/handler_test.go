package grpc

import (
	"github.com/tarmalonchik/speedtest/pkg/api/sdk"
)

var _ sdk.UnitServiceServer = new(UnitSvc)
