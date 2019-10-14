package controller

import (
	"github.com/gin-gonic/gin"
	"gowork/IMService/userServer/model"
)

type UserController struct {
	userInterface model.UserInterface
}

func NewUserController(user model.UserInterface) *UserController {
	return  &UserController{userInterface: user}
}

func (c *UserController) Login(context *gin.Context) {

}

func (c *UserController) Registry(context *gin.Context) {

}