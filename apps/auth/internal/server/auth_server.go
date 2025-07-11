package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/yclw/mys_project/apps/auth/global"
	"github.com/yclw/mys_project/apps/auth/pkg/service"
	"github.com/yclw/mys_project/pkg/common/discovery"
	"github.com/yclw/mys_project/pkg/common/server"
	v1 "github.com/yclw/mys_project/pkg/protobuf/gen/auth/v1"
	"google.golang.org/grpc"
)

var (
	grpcServer   *grpc.Server
	etcdRegister *discovery.Register
)

func InitServer() {
	cfg := global.Cfg
	srv := &http.Server{
		Addr: cfg.Server.Addr,
	}

	// 创建启停服务器
	gracefulSrv := server.NewHttpServer(srv, 10*time.Second)

	// 添加清理函数
	gracefulSrv.AddCleanup(func() error {
		slog.Info("Cleaning up resources...")
		StopGRPCServer()
		return nil
	})

	// 启动服务器
	gracefulSrv.Start()
}

func InitGRPCServer() {
	cfg := global.Cfg

	grpcServer = server.StartGrpcServer(cfg.GrpcServer.Addr, func(s *grpc.Server) {
		v1.RegisterAuthServiceServer(s, service.NewAuthService())
	})

	// 注册grpc服务到etcd
	etcdRegister = discovery.NewRegister(cfg.Etcd.Addrs)
	grpcServerInfo := discovery.Server{
		Name:    cfg.GrpcServer.Name,
		Addr:    cfg.GrpcServer.Addr,
		Version: cfg.GrpcServer.Version,
		Weight:  cfg.GrpcServer.Weight,
	}
	if _, err := etcdRegister.Register(grpcServerInfo, 2); err != nil {
		slog.Error("register etcd error", "error", err)
		return
	}
}

func StopGRPCServer() {
	// 首先停止etcd注册
	etcdRegister.Stop()
	server.StopGrpcServer(grpcServer)
}
