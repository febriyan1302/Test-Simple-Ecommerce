package repository

import (
	"github.com/febriyan1302/catalog-service/internal/model"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

type Cart interface {
	GetCartByUserId(userId int) []model.Cart
	CreateCart(cartData model.CartRequest) (bool, error)
}

func ConstructorCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (repo *CartRepository) GetCartByUserId(userId int) []model.Cart {
	var cart []model.Cart
	repo.db.Preload("Product").Find(&cart, "carts.user_id = ?", userId)
	return cart
}

func (repo *CartRepository) CreateCart(cartData model.CartRequest) (bool, error) {
	dataCart := model.Cart{
		UserId:    cartData.UserId,
		ProductId: cartData.ProductId,
		Amount:    cartData.Amount,
	}

	if err := repo.db.Create(&dataCart).Error; err != nil {
		return false, err
	}

	return true, nil
}
