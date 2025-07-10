package logger

import (
	"strings"
)

// LoggerConfig 从配置文件读取的日志配置
type LoggerConfig struct {
	Level      string `mapstructure:"level" yaml:"level"`
	Format     string `mapstructure:"format" yaml:"format"`
	Output     string `mapstructure:"output" yaml:"output"`
	FilePath   string `mapstructure:"file_path" yaml:"file_path"`
	MaxSize    int64  `mapstructure:"max_size" yaml:"max_size"`
	MaxBackups int    `mapstructure:"max_backups" yaml:"max_backups"`
	MaxAge     int    `mapstructure:"max_age" yaml:"max_age"`
	Compress   bool   `mapstructure:"compress" yaml:"compress"`
}

// ParseLevel 解析日志级别
func ParseLevel(level string) Level {
	switch strings.ToLower(level) {
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "warn", "warning":
		return LevelWarn
	case "error":
		return LevelError
	default:
		return LevelInfo
	}
}

// ParseFormat 解析日志格式
func ParseFormat(format string) Format {
	switch strings.ToLower(format) {
	case "json":
		return FormatJSON
	case "text", "txt":
		return FormatText
	default:
		return FormatJSON
	}
}

// FromLoggerConfig 从LoggerConfig创建Config
func FromLoggerConfig(cfg LoggerConfig) *Config {
	return &Config{
		Level:      ParseLevel(cfg.Level),
		Format:     ParseFormat(cfg.Format),
		Output:     cfg.Output,
		FilePath:   cfg.FilePath,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}
}

// DefaultLoggerConfig 默认的LoggerConfig
func DefaultLoggerConfig() LoggerConfig {
	return LoggerConfig{
		Level:      "info",
		Format:     "json",
		Output:     "console",
		FilePath:   "logs/app.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	}
}
