package grpc

import (
	"context"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tarmalonchik/speedtest/pkg/api/sdk"
)

func (s *BankSvc) GetNodeSpeed(ctx context.Context, req *sdk.GetNodeSpeedRequest) (*sdk.GetNodeSpeedResponse, error) {
	if err := validator.New().Struct(req); err != nil {
		return &sdk.GetNodeSpeedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	inbound, outbound := s.svc.GetNodeSpeed(ctx, req.GetIpAddress())
	return &sdk.GetNodeSpeedResponse{
		InboundSpeed:  inbound,
		OutboundSpeed: outbound,
	}, nil
}
