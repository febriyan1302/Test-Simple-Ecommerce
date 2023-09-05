package service

import (
	"github.com/febriyan1302/catalog-service/internal/model"
	"github.com/febriyan1302/catalog-service/internal/repository"
)

type ProductService struct {
	repository repository.Product
}

type Product interface {
	FindAllProduct() ([]model.ProductResponse, error)
	GetProductDetail(idProduct string) (model.ProductResponse, error)
}

func ConstructProductService(repository repository.Product) *ProductService {
	return &ProductService{repository: repository}
}

func (srv *ProductService) FindAllProduct() ([]model.ProductResponse, error) {
	product, err := srv.repository.GetListAllProduct()
	if err != nil {
		return product, err
	}

	return product, nil
}

func (srv *ProductService) GetProductDetail(idProduct string) (model.ProductResponse, error) {
	product, err := srv.repository.GetProductDetail(idProduct)
	if err != nil {
		return product, err
	}

	return product, nil
}
