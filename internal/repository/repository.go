package repository

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

// RegisterRepository register repository
func RegisterRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}
