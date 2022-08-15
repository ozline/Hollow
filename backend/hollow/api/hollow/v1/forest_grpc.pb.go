// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: hollow/v1/forest.proto

package v1

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

// ForestsClient is the client API for Forests service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ForestsClient interface {
	// Push a Leaf
	Push(ctx context.Context, in *PushLeafRequest, opts ...grpc.CallOption) (*PushLeafReply, error)
	// Get Leafs
	Get(ctx context.Context, in *GetLeafsRequest, opts ...grpc.CallOption) (*GetLeafsReply, error)
}

type forestsClient struct {
	cc grpc.ClientConnInterface
}

func NewForestsClient(cc grpc.ClientConnInterface) ForestsClient {
	return &forestsClient{cc}
}

func (c *forestsClient) Push(ctx context.Context, in *PushLeafRequest, opts ...grpc.CallOption) (*PushLeafReply, error) {
	out := new(PushLeafReply)
	err := c.cc.Invoke(ctx, "/forest.v1.Forests/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forestsClient) Get(ctx context.Context, in *GetLeafsRequest, opts ...grpc.CallOption) (*GetLeafsReply, error) {
	out := new(GetLeafsReply)
	err := c.cc.Invoke(ctx, "/forest.v1.Forests/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForestsServer is the server API for Forests service.
// All implementations must embed UnimplementedForestsServer
// for forward compatibility
type ForestsServer interface {
	// Push a Leaf
	Push(context.Context, *PushLeafRequest) (*PushLeafReply, error)
	// Get Leafs
	Get(context.Context, *GetLeafsRequest) (*GetLeafsReply, error)
	mustEmbedUnimplementedForestsServer()
}

// UnimplementedForestsServer must be embedded to have forward compatible implementations.
type UnimplementedForestsServer struct {
}

func (UnimplementedForestsServer) Push(context.Context, *PushLeafRequest) (*PushLeafReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedForestsServer) Get(context.Context, *GetLeafsRequest) (*GetLeafsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedForestsServer) mustEmbedUnimplementedForestsServer() {}

// UnsafeForestsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ForestsServer will
// result in compilation errors.
type UnsafeForestsServer interface {
	mustEmbedUnimplementedForestsServer()
}

func RegisterForestsServer(s grpc.ServiceRegistrar, srv ForestsServer) {
	s.RegisterService(&Forests_ServiceDesc, srv)
}

func _Forests_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushLeafRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForestsServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/forest.v1.Forests/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForestsServer).Push(ctx, req.(*PushLeafRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forests_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeafsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForestsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/forest.v1.Forests/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForestsServer).Get(ctx, req.(*GetLeafsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Forests_ServiceDesc is the grpc.ServiceDesc for Forests service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Forests_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "forest.v1.Forests",
	HandlerType: (*ForestsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _Forests_Push_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Forests_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hollow/v1/forest.proto",
}
