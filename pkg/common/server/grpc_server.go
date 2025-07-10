package server

import (
	"log/slog"
	"net"

	"github.com/yclw/mys_project/apps/auth/config"
	"github.com/yclw/mys_project/pkg/common/discovery"
	"github.com/yclw/mys_project/pkg/common/global"
	"google.golang.org/grpc"
)

// 启动gRPC服务
func StartGrpc(registerFunc func(s *grpc.Server)) *grpc.Server {
	cfg := global.Cfg.(*config.Config)
	addr := cfg.GrpcServer.Addr
	s := grpc.NewServer()
	registerFunc(s)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		slog.Error("cannot listen", "error", err)
	}
	go func() {
		slog.Info("grpc server started as: %s \n", addr)
		err = s.Serve(lis)
		if err != nil {
			slog.Error("server started error", "error", err)
		}
	}()
	return s
}

// 注册服务到etcd
func RegisterEtcd() {
	cfg := global.Cfg.(*config.Config)

	name := cfg.GrpcServer.Name
	addr := cfg.GrpcServer.Addr
	version := cfg.GrpcServer.Version
	weight := cfg.GrpcServer.Weight
	info := discovery.Server{
		Name:    name,
		Addr:    addr,
		Version: version,
		Weight:  weight,
	}

	etcdAddrs := cfg.Etcd.Addrs
	r := discovery.NewRegister(etcdAddrs, nil) // 注册etcd
	_, err := r.Register(info, 2)
	if err != nil {
		slog.Error("register etcd error", "error", err)
		return
	}
}
