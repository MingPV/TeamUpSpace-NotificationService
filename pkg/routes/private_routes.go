package routes

import (
	// middleware "github.com/MingPV/NotificationService/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterPrivateRoutes(app fiber.Router, db *gorm.DB) {

	// route := app.Group("/api/v1", middleware.JWTMiddleware())

}
