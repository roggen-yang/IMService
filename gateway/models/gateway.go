package models

import (
	"github.com/go-xorm/xorm"
	"github.com/roggen-yang/IMService/gateway/protocol"
)

type GateWayModelInterface interface {
	FindByToken(token string) (*protocol.GateWay, error)
	FindByServerNameTokenAddressTopic(serverName, topic, token, address string) (*protocol.GateWay, error)
	Insert(gateWay *protocol.GateWay) (*protocol.GateWay, error)
	FindByImAddress(imAddress string) ([]*protocol.GateWay, error)
}

type GateWayModel struct {
	mysql *xorm.Engine
}

func NewGateWayModel(mysql *xorm.Engine) *GateWayModel {
	return &GateWayModel{mysql: mysql}
}

func (g *GateWayModel) FindByToken(token string) (*protocol.GateWay, error) {
	gateWay := new(protocol.GateWay)
	_, err := g.mysql.Where("token = ?", token).Get(gateWay)
	if err != nil {
		return nil, err
	}
	return gateWay, nil
}

func (g *GateWayModel) FindByServerNameTokenAddressTopic(serverName, topic, token, address string) (*protocol.GateWay, error) {
	gateWay := new(protocol.GateWay)
	_, err := g.mysql.Where(
		"token = ? and im_address = ? and topic = ? and server_name = ?",
		token,
		address,
		topic,
		serverName).Get(gateWay)
	if err != nil {
		return nil, err
	}
	return gateWay, nil
}

func (g *GateWayModel) Insert(gateWay *protocol.GateWay) (*protocol.GateWay, error) {
	has, err := g.FindByServerNameTokenAddressTopic(gateWay.ServerName, gateWay.Topic, gateWay.Token, gateWay.ImAddress)
	if has != nil && has.Id > 0 && err == nil {
		return has, nil
	}

	_, err = g.mysql.Insert(gateWay)
	if err != nil {
		return nil, err
	}
	return gateWay, nil
}

func (g *GateWayModel) FindByImAddress(imAddress string) ([]*protocol.GateWay, error) {
	gs := []*protocol.GateWay(nil)
	err := g.mysql.Where("im_address = ?", imAddress).Find(&gs)
	if err != nil {
		return nil, err
	}
	return gs, nil
}
