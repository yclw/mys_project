package client

import (
	"log/slog"

	"github.com/yclw/mys_project/apps/api/global"
	"github.com/yclw/mys_project/pkg/common/registrar/etcd"
	v1 "github.com/yclw/mys_project/pkg/protobuf/gen/user/v1"
	"google.golang.org/grpc/resolver"
)

var clients *Clients

// Clients 客户端管理器
type Clients struct {
	UserClient v1.UserServiceClient
}

// Init 初始化所有gRPC客户端
func Init() {
	slog.Info("Initializing gRPC clients...")

	// 初始化 etcd 解析器
	initEtcdResolver()

	clients = &Clients{
		UserClient: initUserClient(),
	}

	slog.Info("All gRPC clients initialized successfully")
}

// initEtcdResolver 初始化 etcd 解析器
func initEtcdResolver() {
	// 获取配置中的 etcd 地址
	cfg := global.Cfg
	etcdAddrs := cfg.Etcd.Addrs

	slog.Info("Initializing etcd resolver", "etcd_addrs", etcdAddrs)

	// 创建并注册 etcd 解析器
	etcdResolver := etcd.NewResolver(etcdAddrs)
	resolver.Register(etcdResolver)

	slog.Info("Etcd resolver initialized successfully")
}
