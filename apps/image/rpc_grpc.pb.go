// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: mcenter/apps/image/pb/rpc.proto

package image

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
	RPC_CreateImage_FullMethodName           = "/mcenter.image.RPC/CreateImage"
	RPC_DeleteImages_FullMethodName          = "/mcenter.image.RPC/DeleteImages"
	RPC_QueryImageByNamespace_FullMethodName = "/mcenter.image.RPC/QueryImageByNamespace"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	// rpc DescribeImage(DescribeImageRequest) returns(Image);
	// rpc QueryImage(QueryImageRequest) returns(ImageSet);
	CreateImage(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*Image, error)
	DeleteImages(ctx context.Context, in *DeleteImagesRequest, opts ...grpc.CallOption) (*ImageSet, error)
	QueryImageByNamespace(ctx context.Context, in *QueryImageByNamespaceRequest, opts ...grpc.CallOption) (*ImageByNamespaceSet, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) CreateImage(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, RPC_CreateImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeleteImages(ctx context.Context, in *DeleteImagesRequest, opts ...grpc.CallOption) (*ImageSet, error) {
	out := new(ImageSet)
	err := c.cc.Invoke(ctx, RPC_DeleteImages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) QueryImageByNamespace(ctx context.Context, in *QueryImageByNamespaceRequest, opts ...grpc.CallOption) (*ImageByNamespaceSet, error) {
	out := new(ImageByNamespaceSet)
	err := c.cc.Invoke(ctx, RPC_QueryImageByNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	// rpc DescribeImage(DescribeImageRequest) returns(Image);
	// rpc QueryImage(QueryImageRequest) returns(ImageSet);
	CreateImage(context.Context, *CreateImageRequest) (*Image, error)
	DeleteImages(context.Context, *DeleteImagesRequest) (*ImageSet, error)
	QueryImageByNamespace(context.Context, *QueryImageByNamespaceRequest) (*ImageByNamespaceSet, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) CreateImage(context.Context, *CreateImageRequest) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImage not implemented")
}
func (UnimplementedRPCServer) DeleteImages(context.Context, *DeleteImagesRequest) (*ImageSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImages not implemented")
}
func (UnimplementedRPCServer) QueryImageByNamespace(context.Context, *QueryImageByNamespaceRequest) (*ImageByNamespaceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryImageByNamespace not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_CreateImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CreateImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CreateImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CreateImage(ctx, req.(*CreateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DeleteImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DeleteImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DeleteImages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DeleteImages(ctx, req.(*DeleteImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_QueryImageByNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryImageByNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryImageByNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QueryImageByNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryImageByNamespace(ctx, req.(*QueryImageByNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mcenter.image.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateImage",
			Handler:    _RPC_CreateImage_Handler,
		},
		{
			MethodName: "DeleteImages",
			Handler:    _RPC_DeleteImages_Handler,
		},
		{
			MethodName: "QueryImageByNamespace",
			Handler:    _RPC_QueryImageByNamespace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mcenter/apps/image/pb/rpc.proto",
}
