package config

import (
	"github.com/micro/go-micro/config"
	commonConfig "github.com/roggen-yang/IMService/common/config"
)

type ImConfig struct {
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
	RabbitMq *commonConfig.RabbitMq
}

type ImRpcConfig struct {
	Version string
	Topic   string
	Server  struct {
		Name      string
		RateLimit int64
	}
	Etcd struct {
		Address  []string
		UserName string
		Password string
	}
	ImServerList []*commonConfig.ImRpc
}

var (
	RpcConf ImRpcConfig
	ImConf  ImConfig
)

func InitRpcConfig(configFile string) error {
	conf := new(ImRpcConfig)
	if err := config.LoadFile(configFile); err != nil {
		return err
	}
	if err := config.Scan(conf); err != nil {
		return err
	}
	RpcConf = *conf
	return nil
}

func InitImConfig(configFile string) error {
	conf := new(ImConfig)
	if err := config.LoadFile(configFile); err != nil {
		return err
	}
	if err := config.Scan(conf); err != nil {
		return err
	}
	ImConf = *conf
	return nil
}
