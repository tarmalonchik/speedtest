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

func (s *BankSvc) Ping(ctx context.Context, req *sdk.PingRequest) (*sdk.PingResponse, error) {
	if err := validator.New().Struct(req); err != nil {
		return &sdk.PingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.svc.UpdateNodeInCache(ctx, req.IpAddress); err != nil {
		logrus.WithError(trace.FuncNameWithError(err)).Errorf("service error")
		return &sdk.PingResponse{}, status.Error(codes.Internal, "service error")
	}
	return &sdk.PingResponse{}, nil
}
