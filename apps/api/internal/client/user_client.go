package client

import (
	"log/slog"

	v1 "github.com/yclw/mys_project/pkg/protobuf/gen/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// User 获取用户服务客户端
func User() v1.UserServiceClient {
	return clients.UserClient
}

// initUserClient 初始化用户服务客户端
func initUserClient() v1.UserServiceClient {
	slog.Info("Connecting to user service...")

	conn, err := grpc.NewClient(
		"etcd:///user",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		slog.Error("Failed to connect to user service", "error", err)
	}

	slog.Info("User service connected successfully")
	return v1.NewUserServiceClient(conn)
}
