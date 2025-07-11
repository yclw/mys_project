package main

import (
	"log/slog"
	"os"

	"github.com/yclw/mys_project/apps/api/config"
	"github.com/yclw/mys_project/apps/api/global"
	"github.com/yclw/mys_project/apps/api/internal/client"
	"github.com/yclw/mys_project/apps/api/internal/server"
	"github.com/yclw/mys_project/apps/api/routes"
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
	cfg := global.Cfg

	slog.Info("Starting API service", "service", cfg.Server.Name)

	// 初始化所有gRPC客户端
	client.Init()

	// 创建HTTP服务器
	router := routes.SetupRouter()
	server.InitServer(router)
}
