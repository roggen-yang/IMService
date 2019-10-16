package rpc

import (
	"flag"
	"github.com/go-acme/lego/log"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/transport/grpc"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/micro/go-plugins/registry/etcdv3"
	imConfig "github.com/roggen-yang/IMService/imServer/config"
	pb "github.com/roggen-yang/IMService/imServer/protos"
	"github.com/roggen-yang/IMService/imServer/rpcserveriml"
	"github.com/roggen-yang/IMService/imServer/server"
)

func main() {
	imFlag := cli.StringFlag{
		Name:  "f",
		Value: "./config/config_rpc.json",
		Usage: "please use xxx -f config_rpc.json",
	}
	configFile := flag.String(imFlag.Name, imFlag.Value, imFlag.Usage)
	flag.Parse()
	err := imConfig.InitRpcConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	etcdRegistry := etcdv3.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = imConfig.RpcConf.Etcd.Address
		})

	service := micro.NewService(
		micro.Name(imConfig.RpcConf.Server.Name),
		micro.Registry(etcdRegistry),
		micro.Version(imConfig.RpcConf.Version),
		micro.Transport(grpc.NewTransport()),
		micro.Flags(imFlag),
	)

	publisherServerMap := make(map[string]*server.RabbitMqBroker)
	for _, item := range imConfig.RpcConf.ImServerList {
		amqbAddress := item.AmqbAddress
		p, err := server.NewRabbitMqBroker(
			item.Topic,
			rabbitmq.NewBroker(func(options *broker.Options) {
				options.Addrs = amqbAddress
			}),
		)
		if err != nil {
			log.Fatal(err)
		}
		publisherServerMap[item.ServerName+item.Topic] = p
	}

	imRpcServer := rpcserveriml.NewImRpcServerIml(publisherServerMap)
	err = pb.RegisterImHandler(service.Server(), imRpcServer)
	if err != nil {
		log.Fatal(err)
	}

	err = service.Run()
	if err != nil {
		log.Fatal(err)
	}
}
