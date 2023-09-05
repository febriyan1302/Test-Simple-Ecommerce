package model

import (
	"time"
)

type Cart struct {
	ID        int       `gorm:"primaryKey" sql:"AUTO_INCREMENT" json:"id"`
	UserId    int       `gorm:"type:int" json:"user_id"`
	ProductId int       `gorm:"type:int" json:"product_id"`
	Amount    int       `gorm:"type:int" json:"amount"`
	CreatedAt time.Time `gorm:"autoCreateTime:true" json:"created_at"`
	Product   Product   `gorm:"foreignKey:ProductId;references:ID" json:"product"`
}

type CartRequest struct {
	UserId    int `json:"user_id"`
	ProductId int `json:"product_id"`
	Amount    int `json:"amount"`
}

type CartData struct {
	ID        int       `json:"id"`
	Amount    int       `json:"amount"`
	Total     int       `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	Product   Product   `json:"product"`
}

type CartResponse struct {
	GrandTotal int        `json:"grand_total"`
	Cart       []CartData `json:"cart"`
}
