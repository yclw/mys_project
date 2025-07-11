package server

import (
	"context"
	"log/slog"
	"net"
	"time"

	"google.golang.org/grpc"
)

// 启动gRPC服务
func StartGrpcServer(addr string, registerFunc func(s *grpc.Server)) *grpc.Server {
	s := grpc.NewServer()
	registerFunc(s)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		slog.Error("cannot listen", "error", err)
	}
	go func() {
		slog.Info("grpc server started", "addr", addr)
		err = s.Serve(lis)
		if err != nil {
			slog.Error("server started error", "error", err)
		}
	}()
	return s
}

func StopGrpcServer(grpcServer *grpc.Server) {
	// 创建一个带超时的上下文，用于优雅停止
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 在goroutine中执行优雅停止
	done := make(chan struct{})
	go func() {
		slog.Info("开始优雅停止 gRPC 服务器...")
		grpcServer.GracefulStop()
		close(done)
	}()

	// 等待优雅停止完成或超时
	select {
	case <-done:
		slog.Info("gRPC 服务器优雅停止完成")
	case <-ctx.Done():
		slog.Warn("优雅停止超时，开始强制停止...")
		grpcServer.Stop()
		slog.Info("gRPC 服务器强制停止完成")
	}
}
