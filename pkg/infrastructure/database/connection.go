package database

import (
	"fmt"
	"log"

	"github.com/febriyan1302/catalog-service/config"
	"github.com/febriyan1302/catalog-service/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

type IConnection interface {
	PGConnect(config *config.Config) (db *gorm.DB, err error)
}

func New() IConnection {
	return &Connection{
		db: &gorm.DB{},
	}
}

// PGConnect Connection to database postgre
func (conn *Connection) PGConnect(config *config.Config) (db *gorm.DB, err error) {
	var (
		dbUser = config.Database.User
		dbPass = config.Database.Password
		dbHost = config.Database.Host
		dbName = config.Database.Name
		dbPort = config.Database.Port
	)

	// dsn
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s`,
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
	)

	// log.Print("dsn:", dsn)
	// db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Connected to database Failed:", err)
		return db, fmt.Errorf("Failed not connect DB: " + err.Error())
	}

	log.Println("Connected to database")

	// Migration
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Cart{})

	// Seeder
	SeedsCategory(db)
	SeedsProduct(db)

	return db, nil
}
