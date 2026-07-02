package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app fiber.Router, db *gorm.DB) {
	api := app.Group("/api")

	productRegisterRoute(api, db)
	stockRegisterRoute(api, db)
}