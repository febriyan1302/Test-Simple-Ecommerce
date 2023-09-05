package service

import (
	"github.com/febriyan1302/catalog-service/internal/model"
	"github.com/febriyan1302/catalog-service/internal/repository"
	"github.com/febriyan1302/catalog-service/pkg/middleware"
)

type AuthService struct {
	repository repository.Auth
}

type Auth interface {
	ValidateAuth(authRequest model.AuthRequest) (authResponse model.AuthResponse, err error)
}

func ConstructAuthService(repository repository.Auth) *AuthService {
	return &AuthService{repository: repository}
}

func (srv *AuthService) ValidateAuth(authRequest model.AuthRequest) (authResponse model.AuthResponse, err error) {

	var response model.AuthResponse

	user, err := srv.repository.FindByCredentials(authRequest.Email, authRequest.Password)
	if err != nil {
		return response, err
	}

	// Generate Token
	token, err := middleware.CreateToken(*user)
	if err != nil {
		return response, err
	}

	response.Token = token
	return response, nil
}
