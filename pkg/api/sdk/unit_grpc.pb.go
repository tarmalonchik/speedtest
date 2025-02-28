// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: unit.proto

package sdk

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UnitService_Measure_FullMethodName = "/sdk.UnitService/Measure"
)

// UnitServiceClient is the client API for UnitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UnitServiceClient interface {
	Measure(ctx context.Context, in *MeasureRequest, opts ...grpc.CallOption) (*MeasureResponse, error)
}

type unitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUnitServiceClient(cc grpc.ClientConnInterface) UnitServiceClient {
	return &unitServiceClient{cc}
}

func (c *unitServiceClient) Measure(ctx context.Context, in *MeasureRequest, opts ...grpc.CallOption) (*MeasureResponse, error) {
	out := new(MeasureResponse)
	err := c.cc.Invoke(ctx, UnitService_Measure_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnitServiceServer is the server API for UnitService service.
// All implementations must embed UnimplementedUnitServiceServer
// for forward compatibility
type UnitServiceServer interface {
	Measure(context.Context, *MeasureRequest) (*MeasureResponse, error)
	mustEmbedUnimplementedUnitServiceServer()
}

// UnimplementedUnitServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUnitServiceServer struct {
}

func (UnimplementedUnitServiceServer) Measure(context.Context, *MeasureRequest) (*MeasureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Measure not implemented")
}
func (UnimplementedUnitServiceServer) mustEmbedUnimplementedUnitServiceServer() {}

// UnsafeUnitServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UnitServiceServer will
// result in compilation errors.
type UnsafeUnitServiceServer interface {
	mustEmbedUnimplementedUnitServiceServer()
}

func RegisterUnitServiceServer(s grpc.ServiceRegistrar, srv UnitServiceServer) {
	s.RegisterService(&UnitService_ServiceDesc, srv)
}

func _UnitService_Measure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeasureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnitServiceServer).Measure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UnitService_Measure_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnitServiceServer).Measure(ctx, req.(*MeasureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UnitService_ServiceDesc is the grpc.ServiceDesc for UnitService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UnitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sdk.UnitService",
	HandlerType: (*UnitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Measure",
			Handler:    _UnitService_Measure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "unit.proto",
}
