package rpcServer

import (
	"github.com/roggen-yang/IMService/common/errors"
	"github.com/roggen-yang/IMService/userServer/model"
	pb "github.com/roggen-yang/IMService/userServer/protos"
	"golang.org/x/net/context"
)

type UserRpcServer struct {
	userModel model.UserModelInterface
}

func NewUserRpcServer(userModel model.UserModelInterface) *UserRpcServer {
	return &UserRpcServer{userModel:userModel}
}

func (u *UserRpcServer) FindByToken(ctx context.Context, req *pb.FindByTokenRequest, rsp *pb.UserResponse)error {
	member, err := u.userModel.FindByToken(req.Token)
	if err != nil {
		return errors.NotFoundUserErr
	}
	rsp.Token = member.Token
	rsp.Id = member.Id
	rsp.Username = member.Username
	rsp.Password = member.Password
	return nil
}

func (u *UserRpcServer) FindById(ctx context.Context, req *pb.FindByIdRequest, rsp *pb.UserResponse) error {
	member, err := u.userModel.FindById(req.Id)
	if err != nil {
		return errors.NotFoundUserErr
	}
	rsp.Token = member.Token
	rsp.Id = member.Id
	rsp.Password = member.Password
	return nil
}
