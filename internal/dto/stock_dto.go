package dto

import "time"

type CreateStockRequest struct {
	ProductID uint   `json:"product_id" validate:"required"`
	Type      string `json:"type" validate:"required,oneof=IN OUT"`
	Quantity  int    `json:"quantity" validate:"required,gt=0"`
	Notes     string `json:"notes"`
}

type StockResponse struct {
	ID        uint            `json:"id"`
	ProductID uint            `json:"product_id"`
	Product   *ProductResponse `json:"product,omitempty"`
	Type      string          `json:"type"`
	Quantity  int             `json:"quantity"`
	Notes     string          `json:"notes"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}