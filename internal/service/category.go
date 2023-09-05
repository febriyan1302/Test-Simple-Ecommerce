package service

import (
	"github.com/febriyan1302/catalog-service/internal/model"
	"github.com/febriyan1302/catalog-service/internal/repository"
)

type CategoryService struct {
	repository repository.Category
}

type Category interface {
	FindAllCategory() ([]model.Category, error)
}

func ConstructorCategoryService(repository repository.Category) *CategoryService {
	return &CategoryService{repository: repository}
}

func (srv *CategoryService) FindAllCategory() ([]model.Category, error) {
	categories, err := srv.repository.GetListAllCategory()
	if err != nil {
		return categories, err
	}
	return categories, nil
}
