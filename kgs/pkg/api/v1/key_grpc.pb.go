// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// KeyServiceClient is the client API for KeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyServiceClient interface {
	GenerateKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GenerateKeyResponse, error)
	DeleteKeys(ctx context.Context, in *DeleteKeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type keyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyServiceClient(cc grpc.ClientConnInterface) KeyServiceClient {
	return &keyServiceClient{cc}
}

func (c *keyServiceClient) GenerateKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GenerateKeyResponse, error) {
	out := new(GenerateKeyResponse)
	err := c.cc.Invoke(ctx, "/key.KeyService/GenerateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) DeleteKeys(ctx context.Context, in *DeleteKeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/key.KeyService/DeleteKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyServiceServer is the server API for KeyService service.
// All implementations must embed UnimplementedKeyServiceServer
// for forward compatibility
type KeyServiceServer interface {
	GenerateKey(context.Context, *emptypb.Empty) (*GenerateKeyResponse, error)
	DeleteKeys(context.Context, *DeleteKeyRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedKeyServiceServer()
}

// UnimplementedKeyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeyServiceServer struct {
}

func (UnimplementedKeyServiceServer) GenerateKey(context.Context, *emptypb.Empty) (*GenerateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKey not implemented")
}
func (UnimplementedKeyServiceServer) DeleteKeys(context.Context, *DeleteKeyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKeys not implemented")
}
func (UnimplementedKeyServiceServer) mustEmbedUnimplementedKeyServiceServer() {}

// UnsafeKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyServiceServer will
// result in compilation errors.
type UnsafeKeyServiceServer interface {
	mustEmbedUnimplementedKeyServiceServer()
}

func RegisterKeyServiceServer(s grpc.ServiceRegistrar, srv KeyServiceServer) {
	s.RegisterService(&KeyService_ServiceDesc, srv)
}

func _KeyService_GenerateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).GenerateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/key.KeyService/GenerateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).GenerateKey(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_DeleteKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).DeleteKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/key.KeyService/DeleteKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).DeleteKeys(ctx, req.(*DeleteKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyService_ServiceDesc is the grpc.ServiceDesc for KeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "key.KeyService",
	HandlerType: (*KeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKey",
			Handler:    _KeyService_GenerateKey_Handler,
		},
		{
			MethodName: "DeleteKeys",
			Handler:    _KeyService_DeleteKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "key.proto",
}
