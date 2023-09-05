package repository

import (
	"github.com/febriyan1302/catalog-service/internal/model"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

type Category interface {
	GetListAllCategory() ([]model.Category, error)
}

func ConstructorCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetListAllCategory() ([]model.Category, error) {
	var categorys []model.Category
	repo.db.Find(&categorys)
	return categorys, nil
}
