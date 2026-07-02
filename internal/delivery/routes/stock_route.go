package routes

import (
	"github.com/dmvsnx/inventory-manegement/internal/delivery/handlers"
	"github.com/dmvsnx/inventory-manegement/internal/repository"
	"github.com/dmvsnx/inventory-manegement/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func stockRegisterRoute(app fiber.Router, db *gorm.DB) {
	stockRepo := repository.NewStockRepository(db)
	productRepo := repository.NewProductRepository(db)
	stockUsecase := usecase.NewStockUsecase(stockRepo, productRepo, db)
	stockHandler := handlers.NewStockHandler(stockUsecase)

	app.Post("/stock", stockHandler.CreateStock)
	app.Get("/stock", stockHandler.GetAllStocks)
	app.Get("/stock/:id", stockHandler.GetStockByID)
	app.Get("/stock/product/:productId", stockHandler.GetStocksByProduct)
}