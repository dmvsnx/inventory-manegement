package model

import (
	"time"

	"gorm.io/gorm"
)

type Type string

const (
	StockIn  Type = "IN"
	StockOut Type = "OUT"
)

func (t Type) IsValid() bool {
	return t == StockIn || t == StockOut
}

type Stock struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	ProductID uint           `gorm:"not null;index" json:"product_id"`
	Product   Product        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product"`
	Type      Type           `gorm:"size:10;not null" json:"type"`
	Quantity  int            `gorm:"not null" json:"quantity"`
	Notes     string         `gorm:"type:text" json:"notes"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
