package repository

import (
	"github.com/febriyan1302/catalog-service/internal/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

type Product interface {
	GetListAllProduct() ([]model.ProductResponse, error)
	GetProductDetail(idProduct string) (model.ProductResponse, error)
}

func ConstructProductRepopsitory(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetListAllProduct() ([]model.ProductResponse, error) {
	var products []model.ProductResponse

	repo.db.Table("products").
		Joins("LEFT JOIN categories on categories.id = products.category_id").
		Select("products.id, products.name, products.description, products.category_id, categories.name as category_name, products.price, products.stock, products.created_at, products.updated_at").
		Find(&products)

	return products, nil
}

func (repo *ProductRepository) GetProductDetail(idProduct string) (model.ProductResponse, error) {
	var product model.ProductResponse
	repo.db.Table("products").
		Joins("LEFT JOIN categories on categories.id = products.category_id").
		Select("products.id, products.name, products.description, products.category_id, categories.name as category_name, products.price, products.stock, products.created_at, products.updated_at").
		Find(&product, idProduct)
	// repo.db.Find(&product, idProduct)

	return product, nil
}
