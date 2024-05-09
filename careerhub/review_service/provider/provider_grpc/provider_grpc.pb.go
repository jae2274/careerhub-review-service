// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: careerhub/review_service/provider/provider_grpc/provider.proto

package provider_grpc

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

// CrawlingTaskGrpcClient is the client API for CrawlingTaskGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrawlingTaskGrpcClient interface {
	AddCrawlingTask(ctx context.Context, in *AddCrawlingTaskRequest, opts ...grpc.CallOption) (*AddCrawlingTaskResponse, error)
}

type crawlingTaskGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewCrawlingTaskGrpcClient(cc grpc.ClientConnInterface) CrawlingTaskGrpcClient {
	return &crawlingTaskGrpcClient{cc}
}

func (c *crawlingTaskGrpcClient) AddCrawlingTask(ctx context.Context, in *AddCrawlingTaskRequest, opts ...grpc.CallOption) (*AddCrawlingTaskResponse, error) {
	out := new(AddCrawlingTaskResponse)
	err := c.cc.Invoke(ctx, "/careerhub.review_service.provider_grpc.CrawlingTaskGrpc/addCrawlingTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrawlingTaskGrpcServer is the server API for CrawlingTaskGrpc service.
// All implementations must embed UnimplementedCrawlingTaskGrpcServer
// for forward compatibility
type CrawlingTaskGrpcServer interface {
	AddCrawlingTask(context.Context, *AddCrawlingTaskRequest) (*AddCrawlingTaskResponse, error)
	mustEmbedUnimplementedCrawlingTaskGrpcServer()
}

// UnimplementedCrawlingTaskGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedCrawlingTaskGrpcServer struct {
}

func (UnimplementedCrawlingTaskGrpcServer) AddCrawlingTask(context.Context, *AddCrawlingTaskRequest) (*AddCrawlingTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCrawlingTask not implemented")
}
func (UnimplementedCrawlingTaskGrpcServer) mustEmbedUnimplementedCrawlingTaskGrpcServer() {}

// UnsafeCrawlingTaskGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrawlingTaskGrpcServer will
// result in compilation errors.
type UnsafeCrawlingTaskGrpcServer interface {
	mustEmbedUnimplementedCrawlingTaskGrpcServer()
}

func RegisterCrawlingTaskGrpcServer(s grpc.ServiceRegistrar, srv CrawlingTaskGrpcServer) {
	s.RegisterService(&CrawlingTaskGrpc_ServiceDesc, srv)
}

func _CrawlingTaskGrpc_AddCrawlingTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCrawlingTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlingTaskGrpcServer).AddCrawlingTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.review_service.provider_grpc.CrawlingTaskGrpc/addCrawlingTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlingTaskGrpcServer).AddCrawlingTask(ctx, req.(*AddCrawlingTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CrawlingTaskGrpc_ServiceDesc is the grpc.ServiceDesc for CrawlingTaskGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrawlingTaskGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "careerhub.review_service.provider_grpc.CrawlingTaskGrpc",
	HandlerType: (*CrawlingTaskGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addCrawlingTask",
			Handler:    _CrawlingTaskGrpc_AddCrawlingTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "careerhub/review_service/provider/provider_grpc/provider.proto",
}
