// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/ad.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AdService_CreateAd_FullMethodName = "/ad.AdService/CreateAd"
	AdService_ReadAd_FullMethodName   = "/ad.AdService/ReadAd"
	AdService_ServeAd_FullMethodName  = "/ad.AdService/ServeAd"
)

// AdServiceClient is the client API for AdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdServiceClient interface {
	CreateAd(ctx context.Context, in *CreateAdRequest, opts ...grpc.CallOption) (*CreateAdResponse, error)
	ReadAd(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*AdResponse, error)
	ServeAd(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*AdResponse, error)
}

type adServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdServiceClient(cc grpc.ClientConnInterface) AdServiceClient {
	return &adServiceClient{cc}
}

func (c *adServiceClient) CreateAd(ctx context.Context, in *CreateAdRequest, opts ...grpc.CallOption) (*CreateAdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAdResponse)
	err := c.cc.Invoke(ctx, AdService_CreateAd_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) ReadAd(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*AdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdResponse)
	err := c.cc.Invoke(ctx, AdService_ReadAd_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) ServeAd(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*AdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdResponse)
	err := c.cc.Invoke(ctx, AdService_ServeAd_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdServiceServer is the server API for AdService service.
// All implementations must embed UnimplementedAdServiceServer
// for forward compatibility.
type AdServiceServer interface {
	CreateAd(context.Context, *CreateAdRequest) (*CreateAdResponse, error)
	ReadAd(context.Context, *AdRequest) (*AdResponse, error)
	ServeAd(context.Context, *AdRequest) (*AdResponse, error)
	mustEmbedUnimplementedAdServiceServer()
}

// UnimplementedAdServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdServiceServer struct{}

func (UnimplementedAdServiceServer) CreateAd(context.Context, *CreateAdRequest) (*CreateAdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAd not implemented")
}
func (UnimplementedAdServiceServer) ReadAd(context.Context, *AdRequest) (*AdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAd not implemented")
}
func (UnimplementedAdServiceServer) ServeAd(context.Context, *AdRequest) (*AdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServeAd not implemented")
}
func (UnimplementedAdServiceServer) mustEmbedUnimplementedAdServiceServer() {}
func (UnimplementedAdServiceServer) testEmbeddedByValue()                   {}

// UnsafeAdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdServiceServer will
// result in compilation errors.
type UnsafeAdServiceServer interface {
	mustEmbedUnimplementedAdServiceServer()
}

func RegisterAdServiceServer(s grpc.ServiceRegistrar, srv AdServiceServer) {
	// If the following call pancis, it indicates UnimplementedAdServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdService_ServiceDesc, srv)
}

func _AdService_CreateAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).CreateAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdService_CreateAd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).CreateAd(ctx, req.(*CreateAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_ReadAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).ReadAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdService_ReadAd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).ReadAd(ctx, req.(*AdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_ServeAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).ServeAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdService_ServeAd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).ServeAd(ctx, req.(*AdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdService_ServiceDesc is the grpc.ServiceDesc for AdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ad.AdService",
	HandlerType: (*AdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAd",
			Handler:    _AdService_CreateAd_Handler,
		},
		{
			MethodName: "ReadAd",
			Handler:    _AdService_ReadAd_Handler,
		},
		{
			MethodName: "ServeAd",
			Handler:    _AdService_ServeAd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ad.proto",
}
