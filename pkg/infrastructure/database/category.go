package database

import (
	"log"

	"github.com/febriyan1302/catalog-service/internal/model"
	"gorm.io/gorm"
)

func SeedsCategory(db *gorm.DB) {
	var count int64
	db.Model(&model.Category{}).Count(&count)

	if count == 0 {
		seedCategory := []model.Category{
			{ID: 1, Name: "Daging"},
			{ID: 2, Name: "Sayur"},
			{ID: 3, Name: "Buah"},
			{ID: 4, Name: "Product Kecantikan"},
			{ID: 5, Name: "Perawatan Tubuh"},
			{ID: 6, Name: "Kebersihan Rumah"},
			{ID: 7, Name: "Bahan Pangan"},
			{ID: 8, Name: "Keperluan Memasak"},
			{ID: 9, Name: "Makanan Beku"},
		}

		for _, category := range seedCategory {
			db.Create(&category)
		}

		log.Println("Seeds Category Success !")
	} else {
		log.Println("Category data exist, abort seeds category.")
	}
}
