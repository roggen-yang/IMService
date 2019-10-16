package protocol

type SendMsgRequest struct {
	FromToken     string `json:"fromToken"`
	ToToken       string `json:"toToken"`
	Body          string `json:"body"`
	Timestamp     int64  `json:"timestamp"`
	RemoteAddress string `json:"remoteAddress"`
}

type LoginRequest struct {
	Token string `json:"token"`
}
