package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	// Order
	orderHandler "github.com/MingPV/NotificationService/internal/order/handler/rest"
	orderRepository "github.com/MingPV/NotificationService/internal/order/repository"
	orderUseCase "github.com/MingPV/NotificationService/internal/order/usecase"

	// Notification
	notificationHandler "github.com/MingPV/NotificationService/internal/notification/handler/rest"
	notificationRepository "github.com/MingPV/NotificationService/internal/notification/repository"
	notificationUseCase "github.com/MingPV/NotificationService/internal/notification/usecase"
)

func RegisterPublicRoutes(app fiber.Router, db *gorm.DB) {

	api := app.Group("/api/v1")

	// === Dependency Wiring ===

	// Order
	orderRepo := orderRepository.NewGormOrderRepository(db)
	orderService := orderUseCase.NewOrderService(orderRepo)
	orderHandler := orderHandler.NewHttpOrderHandler(orderService)

	// Notification
	notificationRepo := notificationRepository.NewGormNotificationRepository(db)
	notificationService := notificationUseCase.NewNotificationService(notificationRepo)
	notificationHandler := notificationHandler.NewHttpNotificationHandler(notificationService)

	// === Public Routes ===

	// Order routes
	orderGroup := api.Group("/orders")
	orderGroup.Get("/", orderHandler.FindAllOrders)
	orderGroup.Get("/:id", orderHandler.FindOrderByID)
	orderGroup.Post("/", orderHandler.CreateOrder)
	orderGroup.Patch("/:id", orderHandler.PatchOrder)
	orderGroup.Delete("/:id", orderHandler.DeleteOrder)

	// Notification routes
	notificationGroup := api.Group("/notifications")
	notificationGroup.Get("/", notificationHandler.FindAllNotifications)
	notificationGroup.Get("/:id", notificationHandler.FindNotificationByID)
	notificationGroup.Post("/", notificationHandler.CreateNotification)
	notificationGroup.Patch("/:id", notificationHandler.PatchNotification)
	notificationGroup.Delete("/:id", notificationHandler.DeleteNotification)

}
