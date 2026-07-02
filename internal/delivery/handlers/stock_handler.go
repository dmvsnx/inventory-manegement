package handlers

import (
	"strconv"

	"github.com/dmvsnx/inventory-manegement/internal/dto"
	"github.com/dmvsnx/inventory-manegement/internal/usecase"
	"github.com/dmvsnx/inventory-manegement/internal/utils/services"
	"github.com/gofiber/fiber/v2"
)

type StockHandler struct {
	stockUsecase usecase.StockUsecase
	validator    *services.ValidatorService
}

func NewStockHandler(su usecase.StockUsecase) *StockHandler {
	return &StockHandler{
		stockUsecase: su,
		validator:    services.NewValidatorService(),
	}
}

func (h *StockHandler) CreateStock(c *fiber.Ctx) error {
	var req dto.CreateStockRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resp, err := h.stockUsecase.Create(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "stock movement created successfully",
		"data":    resp,
	})
}

func (h *StockHandler) GetAllStocks(c *fiber.Ctx) error {
	stockType := c.Query("type")
	start := c.Query("start")
	end := c.Query("end")

	if stockType != "" {
		stocks, err := h.stockUsecase.FindByType(stockType)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map{"data": stocks})
	}

	if start != "" && end != "" {
		stocks, err := h.stockUsecase.FindByDateRange(start, end)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map{"data": stocks})
	}

	stocks, err := h.stockUsecase.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch stock movements",
		})
	}

	return c.JSON(fiber.Map{
		"data": stocks,
	})
}

func (h *StockHandler) GetStockByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid stock id",
		})
	}

	stock, err := h.stockUsecase.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch stock movement",
		})
	}
	if stock == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "stock movement not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": stock,
	})
}

func (h *StockHandler) GetStocksByProduct(c *fiber.Ctx) error {
	productID, err := strconv.ParseUint(c.Params("productId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid product id",
		})
	}

	stocks, err := h.stockUsecase.FindByProductID(uint(productID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch stock movements",
		})
	}

	return c.JSON(fiber.Map{
		"data": stocks,
	})
}
