package grpc

import (
	"context"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tarmalonchik/speedtest/pkg/api/sdk"
)

func (s *BankSvc) Ping(ctx context.Context, req *sdk.PingRequest) (*sdk.PingResponse, error) {
	if err := validator.New().Struct(req); err != nil {
		return &sdk.PingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	s.svc.Ping(ctx, req.ExternalIpAddress, req.InternalIpAddress, req.IsClient)
	return &sdk.PingResponse{}, nil
}
