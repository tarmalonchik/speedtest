package grpc

import (
	"context"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tarmalonchik/speedtest/pkg/api/sdk"
)

func (s *BankSvc) AddNode(ctx context.Context, req *sdk.AddNodeRequest) (*sdk.AddNodeResponse, error) {
	if err := validator.New().Struct(req); err != nil {
		return &sdk.AddNodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	s.svc.AddNode(ctx, req.ExternalIpAddress, req.InternalIpAddress, req.IsClient)
	return &sdk.AddNodeResponse{}, nil
}
