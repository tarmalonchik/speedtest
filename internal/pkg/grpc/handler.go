package grpc

import (
	"google.golang.org/grpc"
)

type Handler grpc.ServiceDesc

func (h Handler) GetServiceDesc() *grpc.ServiceDesc {
	sd := grpc.ServiceDesc(h)
	return &sd
}
