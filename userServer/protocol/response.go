package protocol

type LoginResponse struct {
	Token       string `json:"token"`
	AccessToken string `json:"accessToken"`
	ExpireAt    int64  `json:"expireAt"`
	TimeStamp   int64  `json:"timeStamp"`
}

type RegisterResponse struct {
}
