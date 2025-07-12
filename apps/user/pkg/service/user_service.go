package service

import (
	"context"
	"log/slog"

	v1 "github.com/yclw/mys_project/pkg/protobuf/gen/user/v1"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Ping(ctx context.Context, req *v1.PingRequest) (*v1.PongResponse, error) {
	slog.Info("处理ping请求")
	return &v1.PongResponse{
		Message: "pong",
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	slog.Info("处理登录请求")
	return &v1.LoginResponse{}, nil
}

func (s *UserService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	slog.Info("处理注册请求")
	return &v1.RegisterResponse{}, nil
}

func (s *UserService) SendVerificationCode(ctx context.Context, req *v1.SendVerificationCodeRequest) (*v1.SendVerificationCodeResponse, error) {
	slog.Info("处理发送验证码请求", "email", req.Email)
	return &v1.SendVerificationCodeResponse{}, nil
}
