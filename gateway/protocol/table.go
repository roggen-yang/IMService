package protocol

import "time"

type GateWay struct {
	Id         int64
	Token      string    `json:"token" xorm:"varchar(11) notnull 'token'"`
	ImAddress  string    `json:"imAddress" xorm:"varchar(60) notnull 'im_address'"`
	ServerName string    `json:"server_name" xorm:"varchar(60) notnull 'server_name'"`
	Topic      string    `json:"topic" xorm:"varchar(60) notnull 'topic'"`
	CreateTime time.Time `json:"createTime" xorm:"DateTime 'create_time'"`
	UpdateTime time.Time `json:"updateTime" xorm:"DateTime 'update_time'"`
}

func (g *GateWay) TableName() string {
	return "gateway"
}
