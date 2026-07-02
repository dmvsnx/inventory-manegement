package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	SKU          string         `gorm:"size:50;uniqueIndex;not null" json:"sku"`
	Name         string         `gorm:"size:255;not null" json:"name"`
	Description  string         `gorm:"type:text" json:"description"`
	Category     string         `gorm:"size:100" json:"category"`
	Unit         string         `gorm:"size:20;default:pcs" json:"unit"`
	Price        float64        `gorm:"type:decimal(12,2);default:0" json:"price"`
	Stock        int            `gorm:"default:0" json:"stock"`
	MinimumStock int            `gorm:"default:0" json:"minimum_stock"`
	Stocks       []Stock        `gorm:"foreignKey:ProductID" json:"stocks,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
