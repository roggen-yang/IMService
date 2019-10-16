package im

import (
	"flag"
	"github.com/go-acme/lego/log"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/micro/go-plugins/registry/etcdv3"
	imConfig "github.com/roggen-yang/IMService/imServer/config"
	"github.com/roggen-yang/IMService/imServer/server"
)

func main() {
	imFlag := cli.StringFlag{
		Name:  "f",
		Value: "./config/config_im.json",
		Usage: "please use xxx -f config_im.json",
	}
	configFile := flag.String(imFlag.Name, imFlag.Value, imFlag.Usage)
	flag.Parse()
	err := imConfig.InitImConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	etcdRegistry := etcdv3.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = imConfig.ImConf.Etcd.Address
		})
	rabbitMqRegistry := rabbitmq.NewBroker(func(options *broker.Options) {
		options.Addrs = imConfig.ImConf.RabbitMq.Address
	})
	service := micro.NewService(
		micro.Name(imConfig.ImConf.Server.Name),
		micro.Registry(etcdRegistry),
		micro.Version(imConfig.ImConf.Version),
		micro.Flags(imFlag),
	)

	log.Printf("has start listen topic %s\n", imConfig.ImConf.RabbitMq.Topic)
	rabbitMqBroker, err := server.NewRabbitMqBroker(imConfig.ImConf.RabbitMq.Topic, rabbitMqRegistry)
	if err != nil {
		log.Fatal(err)
	}

	imServer, err := server.NewImServer(rabbitMqBroker,
		func(im *server.ImServer) {
			im.Address = imConfig.ImConf.Port
		})
	if err != nil {
		log.Fatal(err)
	}

	go imServer.Subscribe()
	go imServer.Run()

	service.Run()
	if err != nil {
		log.Fatal(err)
	}
}
