package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/yclw/mys_project/apps/api/config"
	"github.com/yclw/mys_project/apps/api/internal/client"
	"github.com/yclw/mys_project/apps/api/routes"
	"github.com/yclw/mys_project/pkg/common/global"
	"github.com/yclw/mys_project/pkg/common/server"
	"github.com/yclw/mys_project/pkg/utils/logger"
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
	cfg := global.Cfg.(*config.Config)

	slog.Info("Starting API service", "service", cfg.Server.Name)

	// 初始化所有gRPC客户端
	client.Init()

	// 初始化路由
	router := routes.SetupRouter()

	// 创建HTTP服务器
	srv := &http.Server{
		Addr:    cfg.Server.Addr,
		Handler: router,
	}

	// 创建启停服务器
	gracefulSrv := server.NewHttpServer(srv, 10*time.Second)

	// 添加清理函数
	gracefulSrv.AddCleanup(func() error {
		slog.Info("Cleaning up resources...")
		return nil
	})

	// 启动服务器
	gracefulSrv.Start()
}
