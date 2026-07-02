package routes

import (
	"github.com/dmvsnx/inventory-manegement/internal/delivery/handlers"
	"github.com/dmvsnx/inventory-manegement/internal/repository"
	"github.com/dmvsnx/inventory-manegement/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func productRegisterRoute(app fiber.Router, db *gorm.DB) {
	productRepo := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handlers.NewProductHandler(productUsecase)

	app.Post("/products", productHandler.CreateProduct)
	app.Get("/products", productHandler.GetAllProducts)
	app.Get("/products/:id", productHandler.GetProductByID)
	app.Put("/products/:id", productHandler.UpdateProduct)
	app.Delete("/products/:id", productHandler.DeleteProduct)
	app.Get("/reports/low-stock", productHandler.GetLowStockReport)
}