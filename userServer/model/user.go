package model

import (
	"github.com/go-xorm/xorm"
	"github.com/roggen-yang/IMService/userServer/protocol"
)

type UserModelInterface interface {
	FindByToken(token string) (*protocol.Members, error)
	FindById(id int64) (*protocol.Members, error)
	InsertMember(member *protocol.Members)(*protocol.Members, error)
	FindByUserName(userName string) (*protocol.Members, error)
}

type MembersModel struct {
	mysql *xorm.Engine
}

func NewMembersModel(mysql *xorm.Engine) *MembersModel {
	return  &MembersModel{ mysql}
}

func (m *MembersModel) FindByToken(token string) (*protocol.Members, error) {
	member := new(protocol.Members)
	if _, err := m.mysql.Where("token=?", token).Get(member); nil != err {
		return nil, err
	}
	return member, nil
}

func (m *MembersModel) FindById(id int64) (*protocol.Members, error) {
	member := new(protocol.Members)
	if _, err := m.mysql.Where("id=?", id).Get(member); err != nil {
		return nil, err
	}
	return member, nil
}

func (m *MembersModel) FindByUserName(userName string) (*protocol.Members, error) {
	member := new(protocol.Members)
	if _, err := m.mysql.Where("username=?", userName).Get(member); err != nil {
		return nil, err
	}
	return member, nil
}

func (m *MembersModel) InsertMember(member *protocol.Members)(*protocol.Members, error) {
	if _, err := m.mysql.Insert(member); err != nil {
		return nil, err
	}
	return member, nil
}