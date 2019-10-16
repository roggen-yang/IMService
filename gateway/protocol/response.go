package protocol

type SendResponse struct {
}

type GetServerAddressResponse struct {
	Address string `json:"address"`
}
