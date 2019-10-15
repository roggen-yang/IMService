package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	userConfig "github.com/roggen-yang/IMService/userServer/config"
	"github.com/roggen-yang/IMService/userServer/controller"
	"github.com/roggen-yang/IMService/userServer/handlers"
	"github.com/roggen-yang/IMService/userServer/model"
	"log"
)

func InitRouter() *gin.Engine {
	engineUser, err := xorm.NewEngine(userConfig.ApiConf.Engine.Name, userConfig.ApiConf.Engine.DataSource)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	userModel := model.NewMembersModel(engineUser)
	userHandler := handlers.NewUserHandler(userModel)
	userController := controller.NewUserController(userHandler)

	userRouterGroup := router.Group("/user")
	{
		userRouterGroup.POST("/login", userController.Login)
		userRouterGroup.POST("/register", userController.Registry)
	}
	return router
}
