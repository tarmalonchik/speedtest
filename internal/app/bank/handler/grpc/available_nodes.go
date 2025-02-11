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

func (s *BankSvc) AvailableNodes(ctx context.Context, req *sdk.AvailableNodesRequest) (*sdk.AvailableNodesResponse, error) {
	if err := validator.New().Struct(req); err != nil {
		return &sdk.AvailableNodesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	ips, err := s.svc.GetAvailableNodes(ctx)
	if err != nil {
		logrus.WithError(trace.FuncNameWithError(err)).Errorf("service error")
		return &sdk.AvailableNodesResponse{}, status.Error(codes.Internal, "service error")
	}
	return &sdk.AvailableNodesResponse{
		Ip: ips,
	}, nil
}
