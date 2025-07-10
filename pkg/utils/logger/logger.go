package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

// Level 日志级别
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

// Format 日志格式
type Format int

const (
	FormatJSON Format = iota
	FormatText
)

// Config 日志配置
type Config struct {
	Level      Level  // 日志级别
	Format     Format // 日志格式
	Output     string // 输出目标 ("console", "file", "both")
	FilePath   string // 日志文件路径
	MaxSize    int64  // 最大文件大小 (MB)
	MaxBackups int    // 最大备份数
	MaxAge     int    // 最大保存天数
	Compress   bool   // 是否压缩
}

// Logger 日志器接口
type Logger interface {
	Debug(ctx context.Context, msg string, args ...any)
	Info(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
	With(args ...any) Logger
	WithContext(ctx context.Context) Logger
}

// logger 实现
type logger struct {
	slog *slog.Logger
}

// DefaultConfig 默认配置
func DefaultConfig() *Config {
	return &Config{
		Level:      LevelInfo,
		Format:     FormatJSON,
		Output:     "console",
		FilePath:   "logs/app.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	}
}

// New 创建新的日志器
func New(config *Config) Logger {
	if config == nil {
		config = DefaultConfig()
	}

	var writers []io.Writer

	// 根据配置选择输出目标
	switch config.Output {
	case "console":
		writers = append(writers, os.Stdout)
	case "file":
		fileWriter := createFileWriter(config)
		writers = append(writers, fileWriter)
	case "both":
		writers = append(writers, os.Stdout)
		fileWriter := createFileWriter(config)
		writers = append(writers, fileWriter)
	default:
		writers = append(writers, os.Stdout)
	}

	// 创建多输出写入器
	var writer io.Writer
	if len(writers) == 1 {
		writer = writers[0]
	} else {
		writer = io.MultiWriter(writers...)
	}

	// 创建处理器
	var handler slog.Handler
	switch config.Format {
	case FormatJSON:
		handler = slog.NewJSONHandler(writer, &slog.HandlerOptions{
			Level:     convertLevel(config.Level),
			AddSource: true,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				// 自定义时间格式
				if a.Key == slog.TimeKey {
					return slog.String("timestamp", a.Value.Time().Format(time.RFC3339))
				}
				return a
			},
		})
	case FormatText:
		handler = slog.NewTextHandler(writer, &slog.HandlerOptions{
			Level:     convertLevel(config.Level),
			AddSource: true,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				// 自定义时间格式
				if a.Key == slog.TimeKey {
					return slog.String("time", a.Value.Time().Format("2006-01-02 15:04:05"))
				}
				return a
			},
		})
	default:
		handler = slog.NewJSONHandler(writer, &slog.HandlerOptions{
			Level:     convertLevel(config.Level),
			AddSource: true,
		})
	}

	return &logger{
		slog: slog.New(handler),
	}
}

// createFileWriter 创建文件写入器
func createFileWriter(config *Config) io.Writer {
	// 创建日志目录
	if config.FilePath != "" {
		dir := filepath.Dir(config.FilePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			// 如果无法创建目录，回退到控制台输出
			return os.Stdout
		}
	}

	// 打开或创建日志文件
	file, err := os.OpenFile(config.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// 如果无法打开文件，回退到控制台输出
		return os.Stdout
	}

	return file
}

// convertLevel 转换日志级别
func convertLevel(level Level) slog.Level {
	switch level {
	case LevelDebug:
		return slog.LevelDebug
	case LevelInfo:
		return slog.LevelInfo
	case LevelWarn:
		return slog.LevelWarn
	case LevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// Debug 记录调试日志
func (l *logger) Debug(ctx context.Context, msg string, args ...any) {
	l.slog.DebugContext(ctx, msg, args...)
}

// Info 记录信息日志
func (l *logger) Info(ctx context.Context, msg string, args ...any) {
	l.slog.InfoContext(ctx, msg, args...)
}

// Warn 记录警告日志
func (l *logger) Warn(ctx context.Context, msg string, args ...any) {
	l.slog.WarnContext(ctx, msg, args...)
}

// Error 记录错误日志
func (l *logger) Error(ctx context.Context, msg string, args ...any) {
	l.slog.ErrorContext(ctx, msg, args...)
}

// With 添加字段
func (l *logger) With(args ...any) Logger {
	return &logger{
		slog: l.slog.With(args...),
	}
}

// WithContext 添加上下文
func (l *logger) WithContext(ctx context.Context) Logger {
	return &logger{
		slog: l.slog.With(slog.String("trace_id", getTraceID(ctx))),
	}
}

// getTraceID 从上下文获取追踪ID
func getTraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	// 尝试从上下文获取trace_id
	if traceID := ctx.Value("trace_id"); traceID != nil {
		if id, ok := traceID.(string); ok {
			return id
		}
	}

	// 如果没有trace_id，可以生成一个
	return generateTraceID()
}

// generateTraceID 生成追踪ID
func generateTraceID() string {
	return time.Now().Format("20060102150405") + randomString(8)
}

// randomString 生成随机字符串
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}

// 全局日志器
var defaultLogger Logger

// Init 初始化全局日志器
func Init(config *Config) {
	defaultLogger = New(config)
}

// GetLogger 获取全局日志器
func GetLogger() Logger {
	if defaultLogger == nil {
		defaultLogger = New(DefaultConfig())
	}
	return defaultLogger
}

// 便捷方法
func Debug(ctx context.Context, msg string, args ...any) {
	GetLogger().Debug(ctx, msg, args...)
}

func Info(ctx context.Context, msg string, args ...any) {
	GetLogger().Info(ctx, msg, args...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	GetLogger().Warn(ctx, msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	GetLogger().Error(ctx, msg, args...)
}

func With(args ...any) Logger {
	return GetLogger().With(args...)
}

func WithContext(ctx context.Context) Logger {
	return GetLogger().WithContext(ctx)
}
