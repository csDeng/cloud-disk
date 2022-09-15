// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package usercenter

import (
	"context"

	"core/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	EmailVerificationRequest  = pb.EmailVerificationRequest
	EmailVerificationResponse = pb.EmailVerificationResponse
	LoginRequest              = pb.LoginRequest
	LoginResponse             = pb.LoginResponse
	UserRegisterRequest       = pb.UserRegisterRequest
	UserRegisterResponse      = pb.UserRegisterResponse

	UserCenter interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Register(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
		EmailVerification(ctx context.Context, in *EmailVerificationRequest, opts ...grpc.CallOption) (*EmailVerificationResponse, error)
	}

	defaultUserCenter struct {
		cli zrpc.Client
	}
)

func NewUserCenter(cli zrpc.Client) UserCenter {
	return &defaultUserCenter{
		cli: cli,
	}
}

func (m *defaultUserCenter) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := pb.NewUserCenterClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserCenter) Register(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	client := pb.NewUserCenterClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserCenter) EmailVerification(ctx context.Context, in *EmailVerificationRequest, opts ...grpc.CallOption) (*EmailVerificationResponse, error) {
	client := pb.NewUserCenterClient(m.cli.Conn())
	return client.EmailVerification(ctx, in, opts...)
}