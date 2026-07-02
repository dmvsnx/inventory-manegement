package repository

import (
	"errors"
	"time"

	"github.com/dmvsnx/inventory-manegement/internal/model"
	"gorm.io/gorm"
)

type StockRepository interface {
	Create(stock *model.Stock) error
	FindAll() ([]model.Stock, error)
	FindByID(id uint) (*model.Stock, error)
	FindByProductID(productID uint) ([]model.Stock, error)
	FindByType(stockType string) ([]model.Stock, error)
	FindByDateRange(start, end time.Time) ([]model.Stock, error)
	Delete(id uint) error
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockRepository{
		db: db,
	}
}

func (r *stockRepository) Create(stock *model.Stock) error {
	return r.db.Create(stock).Error
}

func (r *stockRepository) FindAll() ([]model.Stock, error) {
	var stocks []model.Stock
	err := r.db.Preload("Product").Find(&stocks).Error
	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func (r *stockRepository) FindByID(id uint) (*model.Stock, error) {
	var stock model.Stock
	err := r.db.Preload("Product").First(&stock, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &stock, nil
}

func (r *stockRepository) FindByProductID(productID uint) ([]model.Stock, error) {
	var stocks []model.Stock
	err := r.db.Where("product_id = ?", productID).Order("created_at DESC").Preload("Product").Find(&stocks).Error
	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func (r *stockRepository) FindByType(stockType string) ([]model.Stock, error) {
	var stocks []model.Stock
	err := r.db.Where("type = ?", stockType).Order("created_at DESC").Preload("Product").Find(&stocks).Error
	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func (r *stockRepository) FindByDateRange(start, end time.Time) ([]model.Stock, error) {
	var stocks []model.Stock
	err := r.db.Where("created_at BETWEEN ? AND ?", start, end).Order("created_at DESC").Preload("Product").Find(&stocks).Error
	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func (r *stockRepository) Delete(id uint) error {
	return r.db.Delete(&model.Stock{}, id).Error
}