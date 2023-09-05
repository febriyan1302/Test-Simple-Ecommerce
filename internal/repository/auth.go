package repository

import (
	"errors"

	"github.com/febriyan1302/catalog-service/internal/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

type Auth interface {
	FindByCredentials(email string, password string) (*model.User, error)
}

func ConstructAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (repo *AuthRepository) FindByCredentials(email string, password string) (*model.User, error) {
	var user model.User
	repo.db.Where("email = ?", email).Where("password = ?", password).Find(&user)

	if user.ID > 0 {
		return &user, nil
	}

	return nil, errors.New("user not found or wrong email/password")
}
