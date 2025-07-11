package cache

import (
	"context"
	"time"
)

// Cache 缓存接口
type Cache interface {
	// Set 设置缓存
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	// Get 获取缓存
	Get(ctx context.Context, key string) (string, error)
	// GetBytes 获取缓存字节数组
	GetBytes(ctx context.Context, key string) ([]byte, error)
	// Exists 检查键是否存在
	Exists(ctx context.Context, key string) (bool, error)
	// Delete 删除缓存
	Delete(ctx context.Context, key string) error
	// DeletePattern 根据模式删除缓存
	DeletePattern(ctx context.Context, pattern string) error
	// SetJSON 设置JSON格式缓存
	SetJSON(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	// GetJSON 获取JSON格式缓存
	GetJSON(ctx context.Context, key string, dest interface{}) error
	// Increment 递增
	Increment(ctx context.Context, key string, value int64) (int64, error)
	// Decrement 递减
	Decrement(ctx context.Context, key string, value int64) (int64, error)
	// Expire 设置过期时间
	Expire(ctx context.Context, key string, expiration time.Duration) error
	// TTL 获取剩余过期时间
	TTL(ctx context.Context, key string) (time.Duration, error)
	// Close 关闭连接
	Close() error
}

var (
	engine Cache
)

// CacheType 缓存类型
type CacheType string

const (
	CacheTypeRedis CacheType = "redis"
)

// 缓存类型映射
var cacheTypeMap = map[CacheType]func(dsn string) (Cache, error){
	CacheTypeRedis: NewRedisCache,
}

// Init 初始化缓存
func Init(cacheType CacheType, dsn string) error {
	cache, err := cacheTypeMap[cacheType](dsn)
	if err != nil {
		return err
	}
	engine = cache
	return nil
}

// GetCache 获取缓存实例
func GetCache() Cache {
	return engine
}
