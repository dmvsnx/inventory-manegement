package usecase

import (
	"errors"

	"github.com/dmvsnx/inventory-manegement/internal/dto"
	"github.com/dmvsnx/inventory-manegement/internal/model"
	"github.com/dmvsnx/inventory-manegement/internal/repository"
)

type ProductUsecase interface {
	Create(req *dto.CreateProductRequest) (*dto.ProductResponse, error)
	FindAll() ([]dto.ProductResponse, error)
	FindByID(id uint) (*dto.ProductResponse, error)
	Update(id uint, req *dto.UpdateProductRequest) (*dto.ProductResponse, error)
	Delete(id uint) error
	FindLowStock() ([]dto.ProductResponse, error)
}

type productUsecase struct {
	productRepo repository.ProductRepository
}

func NewProductUsecase(pr repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: pr,
	}
}

func (u *productUsecase) Create(req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	existing, err := u.productRepo.FindBySKU(req.SKU)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("sku already exists")
	}

	if req.Price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}

	unit := req.Unit
	if unit == "" {
		unit = "pcs"
	}

	product := &model.Product{
		SKU:          req.SKU,
		Name:         req.Name,
		Description:  req.Description,
		Category:     req.Category,
		Unit:         unit,
		Price:        req.Price,
		MinimumStock: req.MinimumStock,
	}

	if err := u.productRepo.Create(product); err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:           product.ID,
		SKU:          product.SKU,
		Name:         product.Name,
		Description:  product.Description,
		Category:     product.Category,
		Unit:         product.Unit,
		Price:        product.Price,
		Stock:        product.Stock,
		MinimumStock: product.MinimumStock,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}, nil
}

func (u *productUsecase) FindAll() ([]dto.ProductResponse, error) {
	products, err := u.productRepo.FindAll()
	if err != nil {
		return nil, err
	}

	result := make([]dto.ProductResponse, len(products))
	for i, p := range products {
		result[i] = dto.ProductResponse{
			ID:           p.ID,
			SKU:          p.SKU,
			Name:         p.Name,
			Description:  p.Description,
			Category:     p.Category,
			Unit:         p.Unit,
			Price:        p.Price,
			Stock:        p.Stock,
			MinimumStock: p.MinimumStock,
			CreatedAt:    p.CreatedAt,
			UpdatedAt:    p.UpdatedAt,
		}
	}

	return result, nil
}

func (u *productUsecase) FindByID(id uint) (*dto.ProductResponse, error) {
	product, err := u.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, nil
	}

	return &dto.ProductResponse{
		ID:           product.ID,
		SKU:          product.SKU,
		Name:         product.Name,
		Description:  product.Description,
		Category:     product.Category,
		Unit:         product.Unit,
		Price:        product.Price,
		Stock:        product.Stock,
		MinimumStock: product.MinimumStock,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}, nil
}

func (u *productUsecase) Update(id uint, req *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	product, err := u.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	if req.SKU != product.SKU {
		existing, err := u.productRepo.FindBySKU(req.SKU)
		if err != nil {
			return nil, err
		}
		if existing != nil {
			return nil, errors.New("sku already exists")
		}
	}

	if req.Price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}

	product.SKU = req.SKU
	product.Name = req.Name
	product.Description = req.Description
	product.Category = req.Category
	product.Unit = req.Unit
	product.Price = req.Price
	product.MinimumStock = req.MinimumStock

	if err := u.productRepo.Update(product); err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:           product.ID,
		SKU:          product.SKU,
		Name:         product.Name,
		Description:  product.Description,
		Category:     product.Category,
		Unit:         product.Unit,
		Price:        product.Price,
		Stock:        product.Stock,
		MinimumStock: product.MinimumStock,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}, nil
}

func (u *productUsecase) Delete(id uint) error {
	return u.productRepo.Delete(id)
}

func (u *productUsecase) FindLowStock() ([]dto.ProductResponse, error) {
	products, err := u.productRepo.FindLowStock()
	if err != nil {
		return nil, err
	}

	result := make([]dto.ProductResponse, len(products))
	for i, p := range products {
		result[i] = dto.ProductResponse{
			ID:           p.ID,
			SKU:          p.SKU,
			Name:         p.Name,
			Description:  p.Description,
			Category:     p.Category,
			Unit:         p.Unit,
			Price:        p.Price,
			Stock:        p.Stock,
			MinimumStock: p.MinimumStock,
			CreatedAt:    p.CreatedAt,
			UpdatedAt:    p.UpdatedAt,
		}
	}

	return result, nil
}
