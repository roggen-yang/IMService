package handlers

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/roggen-yang/IMService/common/constant"
	"github.com/roggen-yang/IMService/common/errors"
	"github.com/roggen-yang/IMService/userServer/model"
	"github.com/roggen-yang/IMService/userServer/protocol"
	"time"
)

type UserHandlerInterface interface {
	Login(l *protocol.LoginRequest) (*protocol.LoginResponse, error)
	Register(r *protocol.RegisterRequest)(*protocol.RegisterResponse, error)
}

type UserHandler struct {
	userModel model.UserModelInterface
}

func NewUserHandler(userModel model.UserModelInterface) *UserHandler {
	return &UserHandler{userModel:userModel}
}

func (u *UserHandler) Login(l *protocol.LoginRequest) (*protocol.LoginResponse, error) {
	user, err := u.userModel.FindByUserName(l.Username)
	if err != nil {
		return nil, errors.NotFoundUserErr
	}
	if user.Password != fmt.Sprintf("%x", md5.Sum([]byte(l.Password))) {
		return nil, errors.UserNameOrPasswordErr
	}
	expired := time.Now().Add(148*time.Hour).Unix()
	accessToken, err := u.createAccessToken(expired)
	if err != nil {
		return nil, errors.UserNameOrPasswordErr
	}
	return &protocol.LoginResponse{
		Token:       user.Token,
		AccessToken: accessToken,
		ExpireAt:    expired,
		TimeStamp:   time.Now().Unix(),
	}, nil
}

func (u *UserHandler) Register(r *protocol.RegisterRequest)(*protocol.RegisterResponse, error){
	member := &protocol.Members{
		Token:      uuid.Must(uuid.NewUUID()).String(),
		Username:   r.Username,
		Password:   fmt.Sprintf("%x", md5.Sum([]byte(r.Password))),
	}
	if _, err := u.userModel.InsertMember(member); err != nil {
		return nil, errors.CreateMemberErr
	}
	return &protocol.RegisterResponse{}, nil
}

func (u *UserHandler) createAccessToken(expired int64) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: expired,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(constant.UserSignedKey))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
