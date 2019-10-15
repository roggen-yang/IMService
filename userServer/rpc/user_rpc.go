package main

import (
	"flag"
	"github.com/go-acme/lego/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	rl "github.com/juju/ratelimit"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/transport/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"
	"github.com/roggen-yang/IMService/userServer/model"
	pb "github.com/roggen-yang/IMService/userServer/protos"
	"github.com/roggen-yang/IMService/userServer/rpcServer"
	userRpcConfig "github.com/roggen-yang/IMService/userserver/config"
)

func main() {
	userRpcFlag := cli.StringFlag{
		Name:  "f",
		Usage: "please use xxx -f config_rpc.json",
		Value: "./config/config_rpc.json",
	}
	configFile := flag.String(userRpcFlag.Name, userRpcFlag.Value, userRpcFlag.Usage)
	flag.Parse()
	conf := new(userRpcConfig.RpcConfig)
	if err := config.LoadFile(*configFile); err != nil {
		log.Fatal(err)
	}
	if err := config.Scan(conf); err != nil {
		log.Fatal(err)
	}

	engineUser, err := xorm.NewEngine(conf.Engine.Name, conf.Engine.DataSource)
	if err != nil {
		log.Fatal(err)
	}
	etcdRegistry := etcdv3.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = conf.Etcd.Address
		})
	b := rl.NewBucketWithRate(float64(conf.Server.RateLimit), int64(conf.Server.RateLimit))
	service := micro.NewService(
		micro.Name(conf.Server.Name),
		micro.Registry(etcdRegistry),
		micro.Version(conf.Version),
		micro.Transport(grpc.NewTransport()),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(b, false)),
		micro.Flags(userRpcFlag),
	)
	service.Init()

	userModel := model.NewMembersModel(engineUser)
	userRpcServer := rpcServer.NewUserRpcServer(userModel)
	if err := pb.RegisterUserHandler(service.Server(), userRpcServer); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
