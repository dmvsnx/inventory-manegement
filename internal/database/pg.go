package database

import (
	"fmt"
	"log"

	"github.com/dmvsnx/inventory-manegement/internal/config"
	"github.com/dmvsnx/inventory-manegement/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Database connected successfully")

	if err := db.AutoMigrate(
		&model.Product{},
		&model.Stock{},
	); err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
