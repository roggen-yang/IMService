package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcdv3"
	userConfig "github.com/roggen-yang/IMService/userServer/config"
	"github.com/roggen-yang/IMService/userServer/router"
	"log"
)

func main() {
	userRpcFlag := cli.StringFlag{
		Name:  "f",
		Value: "../config/config_api.json",
		Usage: "please use xxx -f config_rpc.json",
	}
	configFile := flag.String(userRpcFlag.Name, userRpcFlag.Value, userRpcFlag.Usage)
	flag.Parse()
	if err := userConfig.InitApiConfig(*configFile); err != nil {
		log.Fatal(err)
	}

	etcdRegistry := etcdv3.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = userConfig.ApiConf.Etcd.Address
		})

	service := web.NewService(
		web.Name(userConfig.ApiConf.Server.Name),
		web.Registry(etcdRegistry),
		web.Version(userConfig.ApiConf.Version),
		web.Flags(userRpcFlag),
		web.Address(userConfig.ApiConf.Port),
	)

	r := router.InitRouter()
	service.Handle("/", r)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
