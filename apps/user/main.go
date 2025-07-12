package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/yclw/mys_project/apps/user/config"
	"github.com/yclw/mys_project/apps/user/global"
	"github.com/yclw/mys_project/apps/user/internal/server"
	"github.com/yclw/mys_project/pkg/common/cache"
	"github.com/yclw/mys_project/pkg/common/database"
	"github.com/yclw/mys_project/pkg/model"
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
	// 从全局变量获取配置和日志
	cfg := global.Cfg
	slog.Info("Starting service", "service", cfg.Server.Name)

	// 初始化数据库
	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Mysql.Username, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Db)
	if err := database.Init(database.DBTypeMysql, mysqlDSN, &model.User{}); err != nil {
		slog.Error("Failed to initialize database", "error", err)
		return
	}

	// 初始化缓存
	redisDSN := fmt.Sprintf("redis://:%s@%s:%d/%d",
		cfg.Redis.Password, cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Db)
	if err := cache.Init(cache.CacheTypeRedis, redisDSN); err != nil {
		slog.Error("Failed to initialize cache", "error", err)
		return
	}

	// 创建启动服务器
	server.InitServer()
}
