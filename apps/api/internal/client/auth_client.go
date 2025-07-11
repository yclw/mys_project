package client

import (
	"log/slog"

	v1 "github.com/yclw/mys_project/pkg/protobuf/gen/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Auth 获取认证服务客户端
func Auth() v1.AuthServiceClient {
	return clients.AuthClient
}

// initAuthClient 初始化认证服务客户端
func initAuthClient() v1.AuthServiceClient {
	slog.Info("Connecting to auth service...")

	conn, err := grpc.NewClient(
		"etcd:///auth",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		slog.Error("Failed to connect to auth service", "error", err)
	}

	slog.Info("Auth service connected successfully")
	return v1.NewAuthServiceClient(conn)
}
