package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server" yaml:"server"`
	Etcd   EtcdConfig   `mapstructure:"etcd" yaml:"etcd"`
	Log    LogConfig    `mapstructure:"log" yaml:"log"`
}

type ServerConfig struct {
	Name string `mapstructure:"name" yaml:"name"`
	Addr string `mapstructure:"addr" yaml:"addr"`
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
