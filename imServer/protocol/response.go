package protocol

type SendMsgResponse struct {
	FromToken     string `json:"fromToken"`
	Body          string `json:"body"`
	RemoteAddress string `json:"remoteAddress"`
}
