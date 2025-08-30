package routes

import (
	userHandler "github.com/MingPV/NotificationService/internal/user/handler/rest"
	userRepository "github.com/MingPV/NotificationService/internal/user/repository"
	userUseCase "github.com/MingPV/NotificationService/internal/user/usecase"
	middleware "github.com/MingPV/NotificationService/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterPrivateRoutes(app fiber.Router, db *gorm.DB) {

	route := app.Group("/api/v1", middleware.JWTMiddleware())

	userRepo := userRepository.NewGormUserRepository(db)
	NotificationService := userUseCase.NewNotificationService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(NotificationService)

	route.Get("/me", userHandler.GetUser)

}
