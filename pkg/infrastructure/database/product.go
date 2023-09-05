package database

import (
	"log"

	"github.com/febriyan1302/catalog-service/internal/model"
	"gorm.io/gorm"
)

func SeedsProduct(db *gorm.DB) {
	var count int64
	db.Model(&model.Product{}).Count(&count)

	if count == 0 {
		seedProduct := []model.Product{
			{Name: "Ikan", Description: "Deskripsi ikan", CategoryId: 1, Price: 50000, Stock: 20},
			{Name: "Ayam", Description: "Deskripsi ayam", CategoryId: 1, Price: 30000, Stock: 40},
			{Name: "Kangkung", Description: "Deskripsi kangkung", CategoryId: 2, Price: 5000, Stock: 10},
			{Name: "Melon", Description: "Deskripsi Melon", CategoryId: 3, Price: 25000, Stock: 15},
			{Name: "Jeruk", Description: "Deskripsi Jeruk", CategoryId: 3, Price: 12000, Stock: 15},
		}

		for _, produk := range seedProduct {
			db.Create(&produk)
		}

		log.Println("Seeds Product Success !")
	} else {
		log.Println("Product data exist, abort seeds product.")
	}
}
