// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: grpc.proto

package proto

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

// WishListServiceClient is the client API for WishListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WishListServiceClient interface {
	GetMunitionInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MunitionInfoResponse, error)
}

type wishListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWishListServiceClient(cc grpc.ClientConnInterface) WishListServiceClient {
	return &wishListServiceClient{cc}
}

func (c *wishListServiceClient) GetMunitionInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MunitionInfoResponse, error) {
	out := new(MunitionInfoResponse)
	err := c.cc.Invoke(ctx, "/grpc.WishListService/GetMunitionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WishListServiceServer is the server API for WishListService service.
// All implementations must embed UnimplementedWishListServiceServer
// for forward compatibility
type WishListServiceServer interface {
	GetMunitionInfo(context.Context, *Empty) (*MunitionInfoResponse, error)
	mustEmbedUnimplementedWishListServiceServer()
}

// UnimplementedWishListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWishListServiceServer struct {
}

func (UnimplementedWishListServiceServer) GetMunitionInfo(context.Context, *Empty) (*MunitionInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMunitionInfo not implemented")
}
func (UnimplementedWishListServiceServer) mustEmbedUnimplementedWishListServiceServer() {}

// UnsafeWishListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WishListServiceServer will
// result in compilation errors.
type UnsafeWishListServiceServer interface {
	mustEmbedUnimplementedWishListServiceServer()
}

func RegisterWishListServiceServer(s grpc.ServiceRegistrar, srv WishListServiceServer) {
	s.RegisterService(&WishListService_ServiceDesc, srv)
}

func _WishListService_GetMunitionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishListServiceServer).GetMunitionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.WishListService/GetMunitionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishListServiceServer).GetMunitionInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// WishListService_ServiceDesc is the grpc.ServiceDesc for WishListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WishListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.WishListService",
	HandlerType: (*WishListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMunitionInfo",
			Handler:    _WishListService_GetMunitionInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}