package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/roggen-yang/IMService/common/response"
	"github.com/roggen-yang/IMService/userServer/handlers"
	"github.com/roggen-yang/IMService/userServer/protocol"
)

type UserController struct {
	userHandlerInterface handlers.UserHandlerInterface
}

func NewUserController(user handlers.UserHandlerInterface) *UserController {
	return  &UserController{userHandlerInterface: user}
}

func (c *UserController) Login(context *gin.Context) {
	r := new(protocol.LoginRequest)
	if err := context.ShouldBindJSON(r); err != nil {
		response.ParamError(context, err)
		return
	}
	res, err := c.userHandlerInterface.Login(r)
	response.HttpResponse(context, res, err)
	return
}

func (c *UserController) Registry(context *gin.Context) {
	r := new(protocol.RegisterRequest)
	if err := context.ShouldBindJSON(r); err != nil {
		response.ParamError(context, err)
		return
	}
	res, err := c.userHandlerInterface.Register(r)
	response.HttpResponse(context, res, err)
	return
}