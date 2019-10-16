package protocol

import "time"

type SendRequest struct {
	FromToken string    `json:"fromToken"  binding:"required"`
	ToToken   string    `json:"toToken"  binding:"required"`
	Body      string    `json:"body"  binding:"required"`
	Timestamp time.Time `json:"timestamp"`
}

type GetServerAddressRequest struct {
	Token string `json:"token" binding:"required"`
}
