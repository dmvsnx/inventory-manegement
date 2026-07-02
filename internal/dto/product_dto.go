package dto

import "time"

type CreateProductRequest struct {
	SKU          string  `json:"sku" validate:"required"`
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description"`
	Category     string  `json:"category"`
	Unit         string  `json:"unit"`
	Price        float64 `json:"price" validate:"required,gt=0"`
	MinimumStock int     `json:"minimum_stock"`
}

type UpdateProductRequest struct {
	SKU          string  `json:"sku" validate:"required"`
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description"`
	Category     string  `json:"category"`
	Unit         string  `json:"unit"`
	Price        float64 `json:"price" validate:"required,gt=0"`
	MinimumStock int     `json:"minimum_stock"`
}

type ProductResponse struct {
	ID           uint      `json:"id"`
	SKU          string    `json:"sku"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Category     string    `json:"category"`
	Unit         string    `json:"unit"`
	Price        float64   `json:"price"`
	Stock        int       `json:"stock"`
	MinimumStock int       `json:"minimum_stock"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
