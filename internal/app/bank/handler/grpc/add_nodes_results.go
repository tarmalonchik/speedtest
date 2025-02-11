package grpc

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tarmalonchik/speedtest/internal/app/bank/svc"
	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
	"github.com/tarmalonchik/speedtest/pkg/api/sdk"
)

func (s *BankSvc) AddNodesResults(ctx context.Context, req *sdk.AddNodesResultsRequest) (*sdk.AddNodesResultsResponse, error) {
	if err := validator.New().Struct(req); err != nil {
		return &sdk.AddNodesResultsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	svcReq := make([]svc.NodeResult, len(req.Items))
	for i := range req.Items {
		svcReq[i].IP = req.Items[i].IpAddress
		svcReq[i].InboundSpeed = req.Items[i].InboundSpeed
		svcReq[i].OutboundSpeed = req.Items[i].OutboundSpeed
		svcReq[i].CreatedAt = time.Unix(req.Items[i].CreatedAt, 0).UTC()
	}

	if err := s.svc.AddNodesResults(ctx, svcReq); err != nil {
		logrus.WithError(trace.FuncNameWithError(err)).Errorf("service error")
		return &sdk.AddNodesResultsResponse{}, status.Error(codes.Internal, "service error")
	}
	return &sdk.AddNodesResultsResponse{}, nil
}
