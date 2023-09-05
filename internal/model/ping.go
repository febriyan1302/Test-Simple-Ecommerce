package model

type Ping struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
