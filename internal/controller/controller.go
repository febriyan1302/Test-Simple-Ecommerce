package controller

import "github.com/febriyan1302/catalog-service/internal/service"

type Controller struct {
	service service.Service
}

func RegisterController(service service.Service) Controller {
	return Controller{
		service: service,
	}
}
