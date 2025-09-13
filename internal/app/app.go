package app

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/MingPV/NotificationService/internal/entities"

	// Order
	GrpcOrderHandler "github.com/MingPV/NotificationService/internal/order/handler/grpc"
	orderRepository "github.com/MingPV/NotificationService/internal/order/repository"
	orderUseCase "github.com/MingPV/NotificationService/internal/order/usecase"
	orderpb "github.com/MingPV/NotificationService/proto/order"

	// Notification
	grpcNotificationHandler "github.com/MingPV/NotificationService/internal/notification/handler/grpc"
	notificationRepository "github.com/MingPV/NotificationService/internal/notification/repository"
	notificationUseCase "github.com/MingPV/NotificationService/internal/notification/usecase"
	notificationpb "github.com/MingPV/NotificationService/proto/notification"

	"github.com/MingPV/NotificationService/pkg/config"
	"github.com/MingPV/NotificationService/pkg/database"
	"github.com/MingPV/NotificationService/pkg/middleware"
	"github.com/MingPV/NotificationService/pkg/routes"
)

// rest
func SetupRestServer(db *gorm.DB, cfg *config.Config) (*fiber.App, error) {
	app := fiber.New()
	middleware.FiberMiddleware(app)
	// comment out Swagger when testing
	// routes.SwaggerRoute(app)
	routes.RegisterPublicRoutes(app, db)
	routes.RegisterPrivateRoutes(app, db)
	routes.RegisterNotFoundRoute(app)
	return app, nil
}

// grpc
func SetupGrpcServer(db *gorm.DB, cfg *config.Config) (*grpc.Server, error) {
	s := grpc.NewServer()

	// --- OrderService ---
	orderRepo := orderRepository.NewGormOrderRepository(db)
	orderService := orderUseCase.NewOrderService(orderRepo)
	orderHandler := GrpcOrderHandler.NewGrpcOrderHandler(orderService)
	orderpb.RegisterOrderServiceServer(s, orderHandler)

	// --- NotificationService ---
	notifRepo := notificationRepository.NewGormNotificationRepository(db)
	notifService := notificationUseCase.NewNotificationService(notifRepo)
	notifHandler := grpcNotificationHandler.NewGrpcNotificationHandler(notifService)
	notificationpb.RegisterNotificationServiceServer(s, notifHandler)

	return s, nil
}

// dependencies
func SetupDependencies(env string) (*gorm.DB, *config.Config, error) {
	cfg := config.LoadConfig(env)

	db, err := database.Connect(cfg.DatabaseDSN)
	if err != nil {
		return nil, nil, err
	}

	if env == "test" {
		db.Migrator().DropTable(&entities.Order{}, &entities.Notification{})
	}
	if err := db.AutoMigrate(&entities.Order{}, &entities.Notification{}); err != nil {
		return nil, nil, err
	}

	return db, cfg, nil
}
