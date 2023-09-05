package service

import (
	"github.com/febriyan1302/catalog-service/internal/model"
	"github.com/febriyan1302/catalog-service/internal/repository"
)

type CartService struct {
	repository repository.Cart
}

type Cart interface {
	CreateCart(cartData model.CartRequest) (bool, error)
	GetCart(userId int) (model.CartResponse, error)
}

func ConstructorCartService(repository repository.Cart) *CartService {
	return &CartService{repository: repository}
}

func (srv *CartService) CreateCart(cartData model.CartRequest) (bool, error) {
	_, err := srv.repository.CreateCart(cartData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (srv *CartService) GetCart(userId int) (model.CartResponse, error) {
	cartRepositoryResponse := srv.repository.GetCartByUserId(userId)
	var grandTotal int
	var cartData []model.CartData
	for _, cart := range cartRepositoryResponse {
		data := model.CartData{
			ID:        cart.ID,
			Amount:    cart.Amount,
			Total:     cart.Product.Price * cart.Amount,
			CreatedAt: cart.CreatedAt,
			Product:   cart.Product,
		}
		grandTotal = grandTotal + data.Total
		cartData = append(cartData, data)
	}

	cartResponse := model.CartResponse{
		GrandTotal: grandTotal,
		Cart:       cartData,
	}

	return cartResponse, nil
}
