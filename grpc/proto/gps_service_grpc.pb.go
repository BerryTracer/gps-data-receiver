// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: grpc/proto/gps_service.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GPSServiceClient is the client API for GPSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GPSServiceClient interface {
	Save(ctx context.Context, in *GPSData, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FindByDeviceID(ctx context.Context, in *FindByDeviceIDRequest, opts ...grpc.CallOption) (*GPSDataList, error)
	FindByUserID(ctx context.Context, in *FindByUserIDRequest, opts ...grpc.CallOption) (*GPSDataList, error)
}

type gPSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGPSServiceClient(cc grpc.ClientConnInterface) GPSServiceClient {
	return &gPSServiceClient{cc}
}

func (c *gPSServiceClient) Save(ctx context.Context, in *GPSData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/GPSService/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPSServiceClient) FindByDeviceID(ctx context.Context, in *FindByDeviceIDRequest, opts ...grpc.CallOption) (*GPSDataList, error) {
	out := new(GPSDataList)
	err := c.cc.Invoke(ctx, "/GPSService/FindByDeviceID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPSServiceClient) FindByUserID(ctx context.Context, in *FindByUserIDRequest, opts ...grpc.CallOption) (*GPSDataList, error) {
	out := new(GPSDataList)
	err := c.cc.Invoke(ctx, "/GPSService/FindByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GPSServiceServer is the server API for GPSService service.
// All implementations must embed UnimplementedGPSServiceServer
// for forward compatibility
type GPSServiceServer interface {
	Save(context.Context, *GPSData) (*emptypb.Empty, error)
	FindByDeviceID(context.Context, *FindByDeviceIDRequest) (*GPSDataList, error)
	FindByUserID(context.Context, *FindByUserIDRequest) (*GPSDataList, error)
	mustEmbedUnimplementedGPSServiceServer()
}

// UnimplementedGPSServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGPSServiceServer struct {
}

func (UnimplementedGPSServiceServer) Save(context.Context, *GPSData) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedGPSServiceServer) FindByDeviceID(context.Context, *FindByDeviceIDRequest) (*GPSDataList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByDeviceID not implemented")
}
func (UnimplementedGPSServiceServer) FindByUserID(context.Context, *FindByUserIDRequest) (*GPSDataList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByUserID not implemented")
}
func (UnimplementedGPSServiceServer) mustEmbedUnimplementedGPSServiceServer() {}

// UnsafeGPSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GPSServiceServer will
// result in compilation errors.
type UnsafeGPSServiceServer interface {
	mustEmbedUnimplementedGPSServiceServer()
}

func RegisterGPSServiceServer(s grpc.ServiceRegistrar, srv GPSServiceServer) {
	s.RegisterService(&GPSService_ServiceDesc, srv)
}

func _GPSService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GPSData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPSServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GPSService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPSServiceServer).Save(ctx, req.(*GPSData))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPSService_FindByDeviceID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByDeviceIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPSServiceServer).FindByDeviceID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GPSService/FindByDeviceID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPSServiceServer).FindByDeviceID(ctx, req.(*FindByDeviceIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPSService_FindByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPSServiceServer).FindByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GPSService/FindByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPSServiceServer).FindByUserID(ctx, req.(*FindByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GPSService_ServiceDesc is the grpc.ServiceDesc for GPSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GPSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GPSService",
	HandlerType: (*GPSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _GPSService_Save_Handler,
		},
		{
			MethodName: "FindByDeviceID",
			Handler:    _GPSService_FindByDeviceID_Handler,
		},
		{
			MethodName: "FindByUserID",
			Handler:    _GPSService_FindByUserID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto/gps_service.proto",
}