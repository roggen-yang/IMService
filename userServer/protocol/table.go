package protocol

import "time"

type Members struct {
	Id         int64
	Token      string    `json:"token" xorm:"varchar(11) notnull 'token'"`
	Username   string    `json:"username" xorm:"varchar(60) notnull 'username'"`
	Password   string    `json:"password" xorm:"varchar(60) notnull 'password'"`
	CreateTime time.Time `json:"createTime" xorm:"DateTime 'create_time'"`
	UpdateTime time.Time `json:"updateTime" xorm:"DateTime 'update_time'"`
}
