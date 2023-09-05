package service

import (
	"github.com/febriyan1302/catalog-service/internal/model"
	"github.com/febriyan1302/catalog-service/internal/repository"
)

type UserService struct {
	repository repository.User
}

type User interface {
	CreateService(userData model.UserRequest) (bool, error)
}

func ConstructUserService(repository repository.User) *UserService {
	return &UserService{repository: repository}
}

func (srv *UserService) CreateService(userData model.UserRequest) (bool, error) {
	_, err := srv.repository.CreateRepository(userData)
	if err != nil {
		return false, err
	}

	return true, nil
}
