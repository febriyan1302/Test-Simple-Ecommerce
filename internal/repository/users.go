package repository

import (
	"github.com/febriyan1302/catalog-service/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type User interface {
	CreateRepository(dataUser model.UserRequest) (bool, error)
}

func ConstructUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateRepository(dataRequest model.UserRequest) (bool, error) {
	dataUser := model.User{
		Name:     dataRequest.Name,
		Email:    dataRequest.Email,
		Password: dataRequest.Password,
	}
	if err := repo.db.Create(&dataUser).Error; err != nil {
		return false, err
	}

	return true, nil
}
