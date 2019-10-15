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

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token       string `json:"token"`
	AccessToken string `json:"accessToken"`
	ExpireAt    int64  `json:"expireAt"`
	TimeStamp   int64  `json:"timeStamp"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
}
