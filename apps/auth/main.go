package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/yclw/mys_project/apps/auth/config"
	"github.com/yclw/mys_project/apps/auth/pkg/service"
	"github.com/yclw/mys_project/pkg/common/global"
	"github.com/yclw/mys_project/pkg/common/server"
	v1 "github.com/yclw/mys_project/pkg/protobuf/gen/auth/v1"
	"github.com/yclw/mys_project/pkg/utils/logger"
	"google.golang.org/grpc"
)

func init() {
	// 初始化配置
	cfg, err := config.InitConfig("./config/config.yaml")
	if err != nil {
		slog.Error("Failed to initialize config", "error", err)
		os.Exit(1)
	}

	// 初始化日志
	if err = logger.InitLogger(cfg.Log.Level); err != nil {
		slog.Error("Failed to initialize logger", "error", err)
		os.Exit(1)
	}

	// 初始化全局配置
	global.Cfg = cfg
}

func main() {
	// 从全局变量获取配置和日志
	cfg := global.Cfg.(*config.Config)

	slog.Info("Starting Auth service", "service", cfg.Server.Name)

	// 初始化gRPC服务
	s := server.StartGrpc(func(s *grpc.Server) {
		v1.RegisterAuthServiceServer(s, service.NewAuthService())
	})

	// 初始化数据库

	// 初始化Redis

	// 初始化etcd
	server.RegisterEtcd()

	// 创建HTTP服务器
	srv := &http.Server{
		Addr: cfg.Server.Addr,
	}

	// 创建启停服务器
	gracefulSrv := server.NewHttpServer(srv, 10*time.Second)

	// 添加清理函数
	gracefulSrv.AddCleanup(func() error {
		slog.Info("Cleaning up resources...")
		s.Stop()
		return nil
	})

	// 启动服务器
	gracefulSrv.Start()
}
