package client

import (
	"log/slog"

	"github.com/yclw/mys_project/apps/api/config"
	"github.com/yclw/mys_project/pkg/common/discovery"
	"github.com/yclw/mys_project/pkg/common/global"
	v1 "github.com/yclw/mys_project/pkg/protobuf/gen/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

var clients *Clients

// Clients 客户端管理器
type Clients struct {
	AuthClient v1.AuthServiceClient
}

// Init 初始化所有gRPC客户端
func Init() {
	slog.Info("Initializing gRPC clients...")

	// 初始化 etcd 解析器
	initEtcdResolver()

	clients = &Clients{
		AuthClient: initAuthClient(),
	}

	slog.Info("All gRPC clients initialized successfully")
}

// Auth 获取认证服务客户端
func Auth() v1.AuthServiceClient {
	return clients.AuthClient
}

// initEtcdResolver 初始化 etcd 解析器
func initEtcdResolver() {
	// 获取配置中的 etcd 地址
	cfg := global.Cfg.(*config.Config)
	etcdAddrs := cfg.Etcd.Addrs

	slog.Info("Initializing etcd resolver", "etcd_addrs", etcdAddrs)

	// 创建并注册 etcd 解析器
	etcdResolver := discovery.NewResolver(etcdAddrs, nil)
	resolver.Register(etcdResolver)

	slog.Info("Etcd resolver initialized successfully")
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
