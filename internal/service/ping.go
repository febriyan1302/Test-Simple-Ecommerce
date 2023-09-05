package service

import (
	"github.com/febriyan1302/catalog-service/internal/model"
)

func (svc *Service) GetPing() model.Ping {
	response := model.Ping{
		Status:  200,
		Message: "PONG",
	}
	return response
}
