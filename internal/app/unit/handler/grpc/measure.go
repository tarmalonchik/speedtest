package grpc

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
	"github.com/tarmalonchik/speedtest/pkg/api/sdk"
)

func (s *UnitSvc) Measure(ctx context.Context, req *sdk.MeasureRequest) (*sdk.MeasureResponse, error) {
	if err := validator.New().Struct(req); err != nil {
		return &sdk.MeasureResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	inbound, outbound, err := s.svc.MeasureSpeed(ctx, req.GetIperf3ServerIp())
	if err != nil {
		logrus.WithError(trace.FuncNameWithError(err)).Errorf("service error")
		return &sdk.MeasureResponse{}, status.Error(codes.Internal, "service error")
	}
	return &sdk.MeasureResponse{
		InboundSpeed:  inbound,
		OutboundSpeed: outbound,
	}, nil
}
