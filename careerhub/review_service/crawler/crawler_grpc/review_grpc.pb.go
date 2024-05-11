// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: careerhub/review_service/crawler/crawler_grpc/review.proto

package crawler_grpc

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

// ReviewGrpcClient is the client API for ReviewGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewGrpcClient interface {
	GetCrawlingTasks(ctx context.Context, in *GetCrawlingTasksRequest, opts ...grpc.CallOption) (*GetCrawlingTasksResponse, error)
	SetScoreNPage(ctx context.Context, in *SetScoreNPageRequest, opts ...grpc.CallOption) (*SetScoreNPageResponse, error)
	SetNotExist(ctx context.Context, in *SetNotExistRequest, opts ...grpc.CallOption) (*SetNotExistResponse, error)
	GetCrawlingPages(ctx context.Context, in *GetCrawlingPagesRequest, opts ...grpc.CallOption) (*GetCrawlingPagesResponse, error)
}

type reviewGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewGrpcClient(cc grpc.ClientConnInterface) ReviewGrpcClient {
	return &reviewGrpcClient{cc}
}

func (c *reviewGrpcClient) GetCrawlingTasks(ctx context.Context, in *GetCrawlingTasksRequest, opts ...grpc.CallOption) (*GetCrawlingTasksResponse, error) {
	out := new(GetCrawlingTasksResponse)
	err := c.cc.Invoke(ctx, "/careerhub.review_service.crawler_grpc.ReviewGrpc/getCrawlingTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewGrpcClient) SetScoreNPage(ctx context.Context, in *SetScoreNPageRequest, opts ...grpc.CallOption) (*SetScoreNPageResponse, error) {
	out := new(SetScoreNPageResponse)
	err := c.cc.Invoke(ctx, "/careerhub.review_service.crawler_grpc.ReviewGrpc/setScoreNPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewGrpcClient) SetNotExist(ctx context.Context, in *SetNotExistRequest, opts ...grpc.CallOption) (*SetNotExistResponse, error) {
	out := new(SetNotExistResponse)
	err := c.cc.Invoke(ctx, "/careerhub.review_service.crawler_grpc.ReviewGrpc/setNotExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewGrpcClient) GetCrawlingPages(ctx context.Context, in *GetCrawlingPagesRequest, opts ...grpc.CallOption) (*GetCrawlingPagesResponse, error) {
	out := new(GetCrawlingPagesResponse)
	err := c.cc.Invoke(ctx, "/careerhub.review_service.crawler_grpc.ReviewGrpc/getCrawlingPages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewGrpcServer is the server API for ReviewGrpc service.
// All implementations must embed UnimplementedReviewGrpcServer
// for forward compatibility
type ReviewGrpcServer interface {
	GetCrawlingTasks(context.Context, *GetCrawlingTasksRequest) (*GetCrawlingTasksResponse, error)
	SetScoreNPage(context.Context, *SetScoreNPageRequest) (*SetScoreNPageResponse, error)
	SetNotExist(context.Context, *SetNotExistRequest) (*SetNotExistResponse, error)
	GetCrawlingPages(context.Context, *GetCrawlingPagesRequest) (*GetCrawlingPagesResponse, error)
	mustEmbedUnimplementedReviewGrpcServer()
}

// UnimplementedReviewGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedReviewGrpcServer struct {
}

func (UnimplementedReviewGrpcServer) GetCrawlingTasks(context.Context, *GetCrawlingTasksRequest) (*GetCrawlingTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCrawlingTasks not implemented")
}
func (UnimplementedReviewGrpcServer) SetScoreNPage(context.Context, *SetScoreNPageRequest) (*SetScoreNPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetScoreNPage not implemented")
}
func (UnimplementedReviewGrpcServer) SetNotExist(context.Context, *SetNotExistRequest) (*SetNotExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNotExist not implemented")
}
func (UnimplementedReviewGrpcServer) GetCrawlingPages(context.Context, *GetCrawlingPagesRequest) (*GetCrawlingPagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCrawlingPages not implemented")
}
func (UnimplementedReviewGrpcServer) mustEmbedUnimplementedReviewGrpcServer() {}

// UnsafeReviewGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewGrpcServer will
// result in compilation errors.
type UnsafeReviewGrpcServer interface {
	mustEmbedUnimplementedReviewGrpcServer()
}

func RegisterReviewGrpcServer(s grpc.ServiceRegistrar, srv ReviewGrpcServer) {
	s.RegisterService(&ReviewGrpc_ServiceDesc, srv)
}

func _ReviewGrpc_GetCrawlingTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCrawlingTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewGrpcServer).GetCrawlingTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.review_service.crawler_grpc.ReviewGrpc/getCrawlingTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewGrpcServer).GetCrawlingTasks(ctx, req.(*GetCrawlingTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewGrpc_SetScoreNPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetScoreNPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewGrpcServer).SetScoreNPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.review_service.crawler_grpc.ReviewGrpc/setScoreNPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewGrpcServer).SetScoreNPage(ctx, req.(*SetScoreNPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewGrpc_SetNotExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNotExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewGrpcServer).SetNotExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.review_service.crawler_grpc.ReviewGrpc/setNotExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewGrpcServer).SetNotExist(ctx, req.(*SetNotExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewGrpc_GetCrawlingPages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCrawlingPagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewGrpcServer).GetCrawlingPages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.review_service.crawler_grpc.ReviewGrpc/getCrawlingPages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewGrpcServer).GetCrawlingPages(ctx, req.(*GetCrawlingPagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReviewGrpc_ServiceDesc is the grpc.ServiceDesc for ReviewGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReviewGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "careerhub.review_service.crawler_grpc.ReviewGrpc",
	HandlerType: (*ReviewGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getCrawlingTasks",
			Handler:    _ReviewGrpc_GetCrawlingTasks_Handler,
		},
		{
			MethodName: "setScoreNPage",
			Handler:    _ReviewGrpc_SetScoreNPage_Handler,
		},
		{
			MethodName: "setNotExist",
			Handler:    _ReviewGrpc_SetNotExist_Handler,
		},
		{
			MethodName: "getCrawlingPages",
			Handler:    _ReviewGrpc_GetCrawlingPages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "careerhub/review_service/crawler/crawler_grpc/review.proto",
}
