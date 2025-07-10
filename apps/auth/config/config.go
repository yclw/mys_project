package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server     ServerConfig     `mapstructure:"server" yaml:"server"`
	GrpcServer GrpcServerConfig `mapstructure:"grpc_server" yaml:"grpc_server"`
	Mysql      MysqlConfig      `mapstructure:"mysql" yaml:"mysql"`
	Redis      RedisConfig      `mapstructure:"redis" yaml:"redis"`
	Etcd       EtcdConfig       `mapstructure:"etcd" yaml:"etcd"`
	Log        LogConfig        `mapstructure:"log" yaml:"log"`
}

type ServerConfig struct {
	Name string `mapstructure:"name" yaml:"name"`
	Addr string `mapstructure:"addr" yaml:"addr"`
}

type GrpcServerConfig struct {
	Name    string `mapstructure:"name" yaml:"name"`
	Addr    string `mapstructure:"addr" yaml:"addr"`
	Version string `mapstructure:"version" yaml:"version"`
	Weight  int    `mapstructure:"weight" yaml:"weight"`
}

type MysqlConfig struct {
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	Db       string `mapstructure:"db" yaml:"db"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	Password string `mapstructure:"password" yaml:"password"`
	Db       int    `mapstructure:"db" yaml:"db"`
}

type EtcdConfig struct {
	Addrs []string `mapstructure:"addrs" yaml:"addrs"`
}

type LogConfig struct {
	Level      string `mapstructure:"level" yaml:"level"`             // debug, info, warn, error
	Format     string `mapstructure:"format" yaml:"format"`           // json, text
	Output     string `mapstructure:"output" yaml:"output"`           // console, file, both
	FilePath   string `mapstructure:"file_path" yaml:"file_path"`     // 日志文件路径
	MaxSize    int64  `mapstructure:"max_size" yaml:"max_size"`       // 最大文件大小 (MB)
	MaxBackups int    `mapstructure:"max_backups" yaml:"max_backups"` // 最大备份数
	MaxAge     int    `mapstructure:"max_age" yaml:"max_age"`         // 最大保存天数
	Compress   bool   `mapstructure:"compress" yaml:"compress"`       // 是否压缩
}

func InitConfig(filePath string) (*Config, error) {
	vip := viper.New()
	vip.SetConfigFile(filePath)
	viper.SetConfigType("yaml")
	if err := vip.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := &Config{}
	if err := vip.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
