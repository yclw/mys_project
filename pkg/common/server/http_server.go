package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// HttpServer 优雅启停服务器的配置
type HttpServer struct {
	Server          *http.Server
	ShutdownTimeout time.Duration
	CleanupFuncs    []func() error
}

// NewHttpServer 创建一个新的优雅启停服务器
func NewHttpServer(server *http.Server, shutdownTimeout time.Duration) *HttpServer {
	return &HttpServer{
		Server:          server,
		ShutdownTimeout: shutdownTimeout,
		CleanupFuncs:    make([]func() error, 0),
	}
}

// AddCleanup 添加清理函数
func (hs *HttpServer) AddCleanup(cleanup func() error) {
	hs.CleanupFuncs = append(hs.CleanupFuncs, cleanup)
}

// Start 启动服务器并处理优雅关闭
func (hs *HttpServer) Start() {
	// 在goroutine中启动服务器
	go func() {
		slog.Info("Starting server on %s", hs.Server.Addr)
		if err := hs.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server: %v", err)
			return
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	hs.waitForShutdown()
}

// waitForShutdown 等待关闭信号并执行优雅关闭
func (hs *HttpServer) waitForShutdown() {
	// 创建一个通道来接收OS信号
	quit := make(chan os.Signal, 1)
	// 监听指定的信号：SIGINT (Ctrl+C) 和 SIGTERM
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞等待信号
	<-quit
	ctx := context.Background()
	slog.Info("Shutdown Server ...")

	// 创建一个超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), hs.ShutdownTimeout)
	defer cancel()

	// 执行清理函数
	hs.runCleanupFuncs()

	// 优雅关闭服务器
	if err := hs.Server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown: %v", err)
	}

	slog.Info("Server exiting")
}

// runCleanupFuncs 执行所有清理函数
func (hs *HttpServer) runCleanupFuncs() {
	for i, cleanup := range hs.CleanupFuncs {
		if err := cleanup(); err != nil {
			slog.Error("Error in cleanup function %d: %v", i, err)
		}
	}
}
