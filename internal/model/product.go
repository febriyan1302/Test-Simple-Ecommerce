package model

import "time"

type Product struct {
	ID          int       `gorm:"primaryKey" sql:"AUTO_INCREMENT" json:"id"`
	Name        string    `gorm:"type:varchar(150)" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	CategoryId  int       `gorm:"type:int" json:"category_id"`
	Price       int       `gorm:"type:int" json:"price"`
	Stock       int       `gorm:"type:int" json:"stock"`
	CreatedAt   time.Time `gorm:"autoCreateTime:true" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:true" json:"updated_at"`
}

type ProductResponse struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CategoryId   int       `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Price        int       `json:"price"`
	Stock        int       `json:"stock"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
