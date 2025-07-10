package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GracefulServer 优雅启停服务器的配置
type GracefulServer struct {
	Server          *http.Server
	ShutdownTimeout time.Duration
	CleanupFuncs    []func() error
}

// NewGracefulServer 创建一个新的优雅启停服务器
func NewGracefulServer(server *http.Server, shutdownTimeout time.Duration) *GracefulServer {
	return &GracefulServer{
		Server:          server,
		ShutdownTimeout: shutdownTimeout,
		CleanupFuncs:    make([]func() error, 0),
	}
}

// AddCleanup 添加清理函数
func (gs *GracefulServer) AddCleanup(cleanup func() error) {
	gs.CleanupFuncs = append(gs.CleanupFuncs, cleanup)
}

// Start 启动服务器并处理优雅关闭
func (gs *GracefulServer) Start() {
	// 在goroutine中启动服务器
	go func() {
		log.Printf("Starting server on %s", gs.Server.Addr)
		if err := gs.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	gs.waitForShutdown()
}

// waitForShutdown 等待关闭信号并执行优雅关闭
func (gs *GracefulServer) waitForShutdown() {
	// 创建一个通道来接收OS信号
	quit := make(chan os.Signal, 1)
	// 监听指定的信号：SIGINT (Ctrl+C) 和 SIGTERM
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞等待信号
	<-quit
	log.Println("Shutdown Server ...")

	// 创建一个超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), gs.ShutdownTimeout)
	defer cancel()

	// 执行清理函数
	gs.runCleanupFuncs()

	// 优雅关闭服务器
	if err := gs.Server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}

// runCleanupFuncs 执行所有清理函数
func (gs *GracefulServer) runCleanupFuncs() {
	for i, cleanup := range gs.CleanupFuncs {
		if err := cleanup(); err != nil {
			log.Printf("Error in cleanup function %d: %v", i, err)
		}
	}
}
