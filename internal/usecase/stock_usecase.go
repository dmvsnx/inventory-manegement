package usecase

import (
	"errors"
	"time"

	"github.com/dmvsnx/inventory-manegement/internal/dto"
	"github.com/dmvsnx/inventory-manegement/internal/model"
	"github.com/dmvsnx/inventory-manegement/internal/repository"
	"gorm.io/gorm"
)

type StockUsecase interface {
	Create(req *dto.CreateStockRequest) (*dto.StockResponse, error)
	FindAll() ([]dto.StockResponse, error)
	FindByID(id uint) (*dto.StockResponse, error)
	FindByProductID(productID uint) ([]dto.StockResponse, error)
	FindByType(stockType string) ([]dto.StockResponse, error)
	FindByDateRange(start, end string) ([]dto.StockResponse, error)
	Delete(id uint) error
}

type stockUsecase struct {
	stockRepo   repository.StockRepository
	productRepo repository.ProductRepository
	db          *gorm.DB
}

func NewStockUsecase(sr repository.StockRepository, pr repository.ProductRepository, db *gorm.DB) StockUsecase {
	return &stockUsecase{
		stockRepo:   sr,
		productRepo: pr,
		db:          db,
	}
}

func (u *stockUsecase) Create(req *dto.CreateStockRequest) (*dto.StockResponse, error) {
	product, err := u.productRepo.FindByID(req.ProductID)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	if req.Type != string(model.StockIn) && req.Type != string(model.StockOut) {
		return nil, errors.New("invalid stock type, must be IN or OUT")
	}

	if req.Quantity <= 0 {
		return nil, errors.New("quantity must be greater than zero")
	}

	if req.Type == string(model.StockOut) && product.Stock < req.Quantity {
		return nil, errors.New("insufficient stock")
	}

	quantity := req.Quantity
	if req.Type == string(model.StockOut) {
		quantity = -req.Quantity
	}

	tx := u.db.Begin()

	stock := &model.Stock{
		ProductID: req.ProductID,
		Type:      model.Type(req.Type),
		Quantity:  req.Quantity,
		Notes:     req.Notes,
	}

	if err := tx.Create(stock).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&model.Product{}).Where("id = ?", req.ProductID).
		UpdateColumn("stock", gorm.Expr("stock + ?", quantity)).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	stock.Product = *product
	stock.Product.Stock += quantity

	return &dto.StockResponse{
		ID:        stock.ID,
		ProductID: stock.ProductID,
		Product: &dto.ProductResponse{
			ID:           stock.Product.ID,
			SKU:          stock.Product.SKU,
			Name:         stock.Product.Name,
			Description:  stock.Product.Description,
			Category:     stock.Product.Category,
			Unit:         stock.Product.Unit,
			Price:        stock.Product.Price,
			Stock:        stock.Product.Stock,
			MinimumStock: stock.Product.MinimumStock,
			CreatedAt:    stock.Product.CreatedAt,
			UpdatedAt:    stock.Product.UpdatedAt,
		},
		Type:      string(stock.Type),
		Quantity:  stock.Quantity,
		Notes:     stock.Notes,
		CreatedAt: stock.CreatedAt,
		UpdatedAt: stock.UpdatedAt,
	}, nil
}

func (u *stockUsecase) FindAll() ([]dto.StockResponse, error) {
	stocks, err := u.stockRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return toStockResponses(stocks), nil
}

func (u *stockUsecase) FindByID(id uint) (*dto.StockResponse, error) {
	stock, err := u.stockRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if stock == nil {
		return nil, nil
	}

	resp := toStockResponse(*stock)
	return &resp, nil
}

func (u *stockUsecase) FindByProductID(productID uint) ([]dto.StockResponse, error) {
	stocks, err := u.stockRepo.FindByProductID(productID)
	if err != nil {
		return nil, err
	}

	return toStockResponses(stocks), nil
}

func (u *stockUsecase) FindByType(stockType string) ([]dto.StockResponse, error) {
	stocks, err := u.stockRepo.FindByType(stockType)
	if err != nil {
		return nil, err
	}

	return toStockResponses(stocks), nil
}

func (u *stockUsecase) FindByDateRange(start, end string) ([]dto.StockResponse, error) {
	startTime, err := time.Parse(time.RFC3339, start)
	if err != nil {
		return nil, errors.New("invalid start date format, use RFC3339")
	}

	endTime, err := time.Parse(time.RFC3339, end)
	if err != nil {
		return nil, errors.New("invalid end date format, use RFC3339")
	}

	stocks, err := u.stockRepo.FindByDateRange(startTime, endTime)
	if err != nil {
		return nil, err
	}

	return toStockResponses(stocks), nil
}

func (u *stockUsecase) Delete(id uint) error {
	return u.stockRepo.Delete(id)
}

func toStockResponse(s model.Stock) dto.StockResponse {
	resp := dto.StockResponse{
		ID:        s.ID,
		ProductID: s.ProductID,
		Type:      string(s.Type),
		Quantity:  s.Quantity,
		Notes:     s.Notes,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}

	if s.Product.ID != 0 {
		resp.Product = &dto.ProductResponse{
			ID:           s.Product.ID,
			SKU:          s.Product.SKU,
			Name:         s.Product.Name,
			Description:  s.Product.Description,
			Category:     s.Product.Category,
			Unit:         s.Product.Unit,
			Price:        s.Product.Price,
			Stock:        s.Product.Stock,
			MinimumStock: s.Product.MinimumStock,
			CreatedAt:    s.Product.CreatedAt,
			UpdatedAt:    s.Product.UpdatedAt,
		}
	}

	return resp
}

func toStockResponses(stocks []model.Stock) []dto.StockResponse {
	result := make([]dto.StockResponse, len(stocks))
	for i, s := range stocks {
		result[i] = toStockResponse(s)
	}
	return result
}
