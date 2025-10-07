package app

import (
	"context"
	"log"

	notificationRepository "github.com/MingPV/NotificationService/internal/notification/repository"
	notificationUseCase "github.com/MingPV/NotificationService/internal/notification/usecase"
	"github.com/MingPV/NotificationService/pkg/database"
	"github.com/MingPV/NotificationService/pkg/mq"
	"github.com/MingPV/NotificationService/utils"
)

func Start() {

	// Setup dependencies: database, and configuration
	db, cfg, err := SetupDependencies("dev")
	if err != nil {
		log.Fatalf("❌ Failed to setup dependencies: %v", err)
	}

	// Setup UseCases
	notifRepo := notificationRepository.NewGormNotificationRepository(db)
	notifService := notificationUseCase.NewNotificationService(notifRepo)

	// Start RabbitMQ consumer
	go mq.StartNotificationConsumer(context.Background(), notifService)

	// Setup REST server
	restApp, err := SetupRestServer(db, cfg)
	if err != nil {
		log.Fatalf("❌ Failed to setup REST server: %v", err)
	}

	// Setup gRPC server
	grpcServer, err := SetupGrpcServer(db, cfg)
	if err != nil {
		log.Fatalf("❌ Failed to setup gRPC server: %v", err)
	}

	// Start REST and gRPC servers
	go utils.StartRestServer(restApp, cfg)
	go utils.StartGrpcServer(grpcServer, cfg)

	// Graceful shutdown listener
	utils.WaitForShutdown([]func(){
		func() {
			log.Println("Shutting down REST server...")
			if err := restApp.Shutdown(); err != nil {
				log.Printf("Error shutting down REST server: %v", err)
			}
		},
		func() {
			log.Println("Shutting down gRPC server...")
			grpcServer.GracefulStop()
		},
		func() {
			if err := database.Close(); err != nil {
				log.Printf("Error closing DB: %v", err)
			}
		},
	})

}
