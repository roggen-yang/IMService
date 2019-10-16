package logic

import (
	"github.com/roggen-yang/IMService/common/config"
	"github.com/roggen-yang/IMService/gateway/models"
	imModel "github.com/roggen-yang/IMService/imServer/rpcserveriml"
	userModel "github.com/roggen-yang/IMService/userServer/model"
)

type GateWayLogic struct {
	userRpcModel  userModel.UserModelInterface
	gateWayModel  models.GateWayModelInterface
	imRpcModel    imModel.ImRpcServerIml
	imAddressList []*config.ImRpc
}

//func NewGtateWayLogic(userRpcModel userModel.UserModelInterface,
//	gateWayModel models.GateWayModelInterface)
