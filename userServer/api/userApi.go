package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/micro/cli"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcdv3"
	userConfig "github.com/roggen-yang/IMService/userServer/config"
	"github.com/roggen-yang/IMService/userServer/controller"
	"github.com/roggen-yang/IMService/userServer/handlers"
	"github.com/roggen-yang/IMService/userServer/model"
	"log"
)

func main() {
	userRpcFlag := cli.StringFlag{
		Name:  "f",
		Value: "./config/config_api.json",
		Usage: "please use xxx -f config_rpc.json",
	}
	configFile := flag.String(userRpcFlag.Name, userRpcFlag.Value, userRpcFlag.Usage)
	flag.Parse()
	conf := new(userConfig.ApiConfig)

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

	service := web.NewService(
		web.Name(conf.Server.Name),
		web.Registry(etcdRegistry),
		web.Version(conf.Version),
		web.Flags(userRpcFlag),
		web.Address(conf.Port),
	)
	router := gin.Default()
	userModel := model.NewMembersModel(engineUser)
	userHandler := handlers.NewUserHandler(userModel)
	userController := controller.NewUserController(userHandler)

	userRouterGroup := router.Group("/user")
	{
		userRouterGroup.POST("/login", userController.Login)
		userRouterGroup.POST("/register", userController.Registry)
	}
	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
