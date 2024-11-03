// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: dapplink/appchain.proto

package appchain

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
	AppChainService_L1StakerRewardsAmount_FullMethodName = "/services.dapplink.appchain.AppChainService/L1StakerRewardsAmount"
	AppChainService_L2StakerRewardsAmount_FullMethodName = "/services.dapplink.appchain.AppChainService/L2StakerRewardsAmount"
)

// AppChainServiceClient is the client API for AppChainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppChainServiceClient interface {
	L1StakerRewardsAmount(ctx context.Context, in *L1StakerRewardsAmountRequest, opts ...grpc.CallOption) (*L1StakerRewardsAmountResponse, error)
	L2StakerRewardsAmount(ctx context.Context, in *L2StakerRewardsAmountRequest, opts ...grpc.CallOption) (*L2StakerRewardsAmountResponse, error)
}

type appChainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppChainServiceClient(cc grpc.ClientConnInterface) AppChainServiceClient {
	return &appChainServiceClient{cc}
}

func (c *appChainServiceClient) L1StakerRewardsAmount(ctx context.Context, in *L1StakerRewardsAmountRequest, opts ...grpc.CallOption) (*L1StakerRewardsAmountResponse, error) {
	out := new(L1StakerRewardsAmountResponse)
	err := c.cc.Invoke(ctx, AppChainService_L1StakerRewardsAmount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appChainServiceClient) L2StakerRewardsAmount(ctx context.Context, in *L2StakerRewardsAmountRequest, opts ...grpc.CallOption) (*L2StakerRewardsAmountResponse, error) {
	out := new(L2StakerRewardsAmountResponse)
	err := c.cc.Invoke(ctx, AppChainService_L2StakerRewardsAmount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppChainServiceServer is the server API for AppChainService service.
// All implementations should embed UnimplementedAppChainServiceServer
// for forward compatibility
type AppChainServiceServer interface {
	L1StakerRewardsAmount(context.Context, *L1StakerRewardsAmountRequest) (*L1StakerRewardsAmountResponse, error)
	L2StakerRewardsAmount(context.Context, *L2StakerRewardsAmountRequest) (*L2StakerRewardsAmountResponse, error)
}

// UnimplementedAppChainServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAppChainServiceServer struct {
}

func (UnimplementedAppChainServiceServer) L1StakerRewardsAmount(context.Context, *L1StakerRewardsAmountRequest) (*L1StakerRewardsAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method L1StakerRewardsAmount not implemented")
}
func (UnimplementedAppChainServiceServer) L2StakerRewardsAmount(context.Context, *L2StakerRewardsAmountRequest) (*L2StakerRewardsAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method L2StakerRewardsAmount not implemented")
}

// UnsafeAppChainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppChainServiceServer will
// result in compilation errors.
type UnsafeAppChainServiceServer interface {
	mustEmbedUnimplementedAppChainServiceServer()
}

func RegisterAppChainServiceServer(s grpc.ServiceRegistrar, srv AppChainServiceServer) {
	s.RegisterService(&AppChainService_ServiceDesc, srv)
}

func _AppChainService_L1StakerRewardsAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(L1StakerRewardsAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppChainServiceServer).L1StakerRewardsAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppChainService_L1StakerRewardsAmount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppChainServiceServer).L1StakerRewardsAmount(ctx, req.(*L1StakerRewardsAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppChainService_L2StakerRewardsAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(L2StakerRewardsAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppChainServiceServer).L2StakerRewardsAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppChainService_L2StakerRewardsAmount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppChainServiceServer).L2StakerRewardsAmount(ctx, req.(*L2StakerRewardsAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppChainService_ServiceDesc is the grpc.ServiceDesc for AppChainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppChainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.dapplink.appchain.AppChainService",
	HandlerType: (*AppChainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "L1StakerRewardsAmount",
			Handler:    _AppChainService_L1StakerRewardsAmount_Handler,
		},
		{
			MethodName: "L2StakerRewardsAmount",
			Handler:    _AppChainService_L2StakerRewardsAmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dapplink/appchain.proto",
}
