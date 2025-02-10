package grpc

import (
	"context"
	"errors"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/tarmalonchik/speedtest/internal/pkg/hashsign"
)

type hashSigner interface {
	SignRequest(body []byte) (hashsign.SignResponse, error)
	CheckRequest(body []byte, opt ...hashsign.CheckOption) (bool, error)
	WithTime(time string) hashsign.CheckOption
	WithHash(hash string) hashsign.CheckOption
}

type Service struct {
	keepalive          serverParameters
	address            string
	grpcServer         *grpc.Server
	serverOptions      []grpc.ServerOption
	streamInterceptors []grpc.StreamServerInterceptor
	unaryInterceptors  []grpc.UnaryServerInterceptor
}

func New(address string, opts ...Options) (*Service, error) {
	gs := &Service{
		address:            address,
		serverOptions:      make([]grpc.ServerOption, 0),
		streamInterceptors: make([]grpc.StreamServerInterceptor, 0),
		unaryInterceptors:  []grpc.UnaryServerInterceptor{},
	}

	for i := range opts {
		opts[i](gs)
	}

	return gs, nil
}

func (gs *Service) Init() error {
	gs.serverOptions = append(gs.serverOptions, grpc.KeepaliveParams(gs.getKeepaliveParams()))
	gs.serverOptions = append(gs.serverOptions, grpc.ChainStreamInterceptor(gs.streamInterceptors...))
	gs.serverOptions = append(gs.serverOptions, grpc.ChainUnaryInterceptor(gs.unaryInterceptors...))

	gs.grpcServer = grpc.NewServer(
		gs.serverOptions...,
	)

	reflection.Register(gs.grpcServer)

	return nil
}

func (gs *Service) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	gs.grpcServer.RegisterService(desc, impl)
}

func (gs *Service) Run(ctx context.Context) error {
	logrus.Infof("grpc server listening: %s", gs.address)

	lis, err := net.Listen("tcp", gs.address)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		logrus.Infof("stopped grpc server")
		gs.grpcServer.GracefulStop()
		err = lis.Close()
		if err != nil && !errors.Is(err, net.ErrClosed) {
			logrus.Errorf("failed to shutdown grpc connect: %s", gs.address)
		}
	}()

	err = gs.grpcServer.Serve(lis)
	if errors.Is(err, grpc.ErrServerStopped) {
		err = nil
	}
	return err
}
