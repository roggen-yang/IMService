package config

import (
	"github.com/micro/go-micro/config"
)

type (
	RpcConfig struct {
		Version string
		Server  struct {
			Name      string
			RateLimit int64
		}
		Etcd struct {
			Address  []string
			UserName string
			Password string
		}
		Engine struct {
			Name       string
			DataSource string
		}
	}

	ApiConfig struct {
		Version string
		Port    string
		Server  struct {
			Name      string
			RateLimit int64
		}
		Etcd struct {
			Address  []string
			UserName string
			Password string
		}
		Engine struct {
			Name       string
			DataSource string
		}
	}
)

var (
	RpcConf RpcConfig
	ApiConf ApiConfig
)

func InitRpcConfig(configFile string) error {
	conf := new(RpcConfig)
	if err := config.LoadFile(configFile); err != nil {
		return err
	}
	if err := config.Scan(conf); err != nil {
		return err
	}
	RpcConf = *conf
	return nil
}

func InitApiConfig(configFile string) error {
	conf := new(ApiConfig)
	if err := config.LoadFile(configFile); err != nil {
		return err
	}
	if err := config.Scan(conf); err != nil {
		return err
	}
	ApiConf = *conf
	return nil
}
