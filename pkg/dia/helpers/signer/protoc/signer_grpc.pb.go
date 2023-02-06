// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: protoc/signer.proto

package signer

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

// SignerClient is the client API for Signer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignerClient interface {
	Sign(ctx context.Context, in *DataToSign, opts ...grpc.CallOption) (*SignedData, error)
}

type signerClient struct {
	cc grpc.ClientConnInterface
}

func NewSignerClient(cc grpc.ClientConnInterface) SignerClient {
	return &signerClient{cc}
}

func (c *signerClient) Sign(ctx context.Context, in *DataToSign, opts ...grpc.CallOption) (*SignedData, error) {
	out := new(SignedData)
	err := c.cc.Invoke(ctx, "/rpc.Signer/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignerServer is the server API for Signer service.
// All implementations must embed UnimplementedSignerServer
// for forward compatibility
type SignerServer interface {
	Sign(context.Context, *DataToSign) (*SignedData, error)
	mustEmbedUnimplementedSignerServer()
}

// UnimplementedSignerServer must be embedded to have forward compatible implementations.
type UnimplementedSignerServer struct {
}

func (UnimplementedSignerServer) Sign(context.Context, *DataToSign) (*SignedData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedSignerServer) mustEmbedUnimplementedSignerServer() {}

// UnsafeSignerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignerServer will
// result in compilation errors.
type UnsafeSignerServer interface {
	mustEmbedUnimplementedSignerServer()
}

func RegisterSignerServer(s grpc.ServiceRegistrar, srv SignerServer) {
	s.RegisterService(&Signer_ServiceDesc, srv)
}

func _Signer_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataToSign)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignerServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Signer/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignerServer).Sign(ctx, req.(*DataToSign))
	}
	return interceptor(ctx, in, info, handler)
}

// Signer_ServiceDesc is the grpc.ServiceDesc for Signer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Signer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Signer",
	HandlerType: (*SignerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sign",
			Handler:    _Signer_Sign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoc/signer.proto",
}