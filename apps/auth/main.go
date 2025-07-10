package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/yclw/mys_project/apps/auth/config"
	"github.com/yclw/mys_project/pkg/common/global"
	"github.com/yclw/mys_project/pkg/common/server"
	"github.com/yclw/mys_project/pkg/utils/logger"
)

func init() {
	// 初始化配置
	cfg, err := config.InitConfig("./config/config.yaml")
	if err != nil {
		log.Fatal("Failed to initialize config:", err)
		return
	}

	// 初始化日志
	logger.Init(&logger.Config{
		Level:      logger.ParseLevel(cfg.Log.Level),
		Format:     logger.ParseFormat(cfg.Log.Format),
		Output:     cfg.Log.Output,
		FilePath:   cfg.Log.FilePath,
		MaxSize:    cfg.Log.MaxSize,
		MaxBackups: cfg.Log.MaxBackups,
		MaxAge:     cfg.Log.MaxAge,
		Compress:   cfg.Log.Compress,
	})

	// 初始化全局配置
	global.Cfg = cfg
	global.Logs = logger.GetLogger()
}

func main() {
	ctx := context.Background()

	// 从全局变量获取配置和日志
	cfg := global.Cfg.(*config.Config)

	global.Logs.Info(ctx, "Starting Auth service", "service", cfg.Server.Name)

	// 初始化gRPC服务

	// 初始化数据库

	// 初始化Redis

	// 初始化etcd

	// 创建HTTP服务器
	srv := &http.Server{
		Addr: cfg.Server.Addr,
	}

	// 创建启停服务器
	gracefulSrv := server.NewGracefulServer(srv, 10*time.Second)

	// 添加清理函数
	gracefulSrv.AddCleanup(func() error {
		global.Logs.Info(ctx, "Cleaning up resources...")
		return nil
	})

	// 启动服务器
	gracefulSrv.Start()
}
