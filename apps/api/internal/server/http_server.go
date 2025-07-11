package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yclw/mys_project/apps/api/global"
	"github.com/yclw/mys_project/pkg/common/server"
)

func InitServer(router *gin.Engine) {
	cfg := global.Cfg
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
