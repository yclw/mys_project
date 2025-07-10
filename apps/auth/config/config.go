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
	Level string `mapstructure:"level" yaml:"level"` // debug, info, warn, error
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
