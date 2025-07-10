package logger

import (
	"log/slog"
	"strings"
)

// 初始化
func InitLogger(level string) error {
	// //使用json格式
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	// 	AddSource: true,
	// 	Level:     LogLevel(level),
	// }))
	// slog.SetDefault(logger)
	return nil
}

// 获得日志等级
func LogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	}
	return slog.LevelInfo
}
