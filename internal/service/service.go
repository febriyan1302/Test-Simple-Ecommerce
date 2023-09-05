package service

import (
	"github.com/febriyan1302/catalog-service/internal/model"
	"github.com/febriyan1302/catalog-service/internal/repository"
)

type Service struct {
	repository repository.Repository
}

// RegisterService register service
func RegisterService(repo repository.Repository) Service {
	return Service{
		repository: repo,
	}
}

// IService register the service to expose public access
type IService interface {
	GetPing() model.Ping
}

// New call the instance service
func New() IService {
	return &Service{}
}
