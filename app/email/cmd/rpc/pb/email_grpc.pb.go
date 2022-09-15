// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: email.proto

package pb

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

// EmailCenterClient is the client API for EmailCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailCenterClient interface {
	SendCode(ctx context.Context, in *SendCodeRequest, opts ...grpc.CallOption) (*SendCodeResponse, error)
}

type emailCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailCenterClient(cc grpc.ClientConnInterface) EmailCenterClient {
	return &emailCenterClient{cc}
}

func (c *emailCenterClient) SendCode(ctx context.Context, in *SendCodeRequest, opts ...grpc.CallOption) (*SendCodeResponse, error) {
	out := new(SendCodeResponse)
	err := c.cc.Invoke(ctx, "/pb.EmailCenter/SendCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailCenterServer is the server API for EmailCenter service.
// All implementations must embed UnimplementedEmailCenterServer
// for forward compatibility
type EmailCenterServer interface {
	SendCode(context.Context, *SendCodeRequest) (*SendCodeResponse, error)
	mustEmbedUnimplementedEmailCenterServer()
}

// UnimplementedEmailCenterServer must be embedded to have forward compatible implementations.
type UnimplementedEmailCenterServer struct {
}

func (UnimplementedEmailCenterServer) SendCode(context.Context, *SendCodeRequest) (*SendCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCode not implemented")
}
func (UnimplementedEmailCenterServer) mustEmbedUnimplementedEmailCenterServer() {}

// UnsafeEmailCenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailCenterServer will
// result in compilation errors.
type UnsafeEmailCenterServer interface {
	mustEmbedUnimplementedEmailCenterServer()
}

func RegisterEmailCenterServer(s grpc.ServiceRegistrar, srv EmailCenterServer) {
	s.RegisterService(&EmailCenter_ServiceDesc, srv)
}

func _EmailCenter_SendCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailCenterServer).SendCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EmailCenter/SendCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailCenterServer).SendCode(ctx, req.(*SendCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailCenter_ServiceDesc is the grpc.ServiceDesc for EmailCenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailCenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EmailCenter",
	HandlerType: (*EmailCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCode",
			Handler:    _EmailCenter_SendCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email.proto",
}