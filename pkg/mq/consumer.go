package mq

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/MingPV/NotificationService/internal/entities"
	"github.com/MingPV/NotificationService/internal/notification/usecase"
	"github.com/streadway/amqp"
)

type MessageEnvelope struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// StartNotificationConsumer connect RabbitMQ และ consume messages
func StartNotificationConsumer(ctx context.Context, notificationUC usecase.NotificationUseCase) {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@host.docker.internal:5672/"
	}

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("Failed to connect RabbitMQ: %v", err)
	}
	// defer conn.Close() // close when application shutdown
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}

	q, err := ch.QueueDeclare(
		"notifications",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	log.Println("[*] Notification consumer started, waiting for messages...")

	// go func() {
	// 	for d := range msgs {
	// 		var event entities.EventCreatedEvent
	// 		if err := json.Unmarshal(d.Body, &event); err != nil {
	// 			log.Println("Failed to decode message:", err)
	// 			continue
	// 		}

	// 		// Call UseCase
	// 		notificationUC.HandleEventCreatedEvent(context.Background(), &event)
	// 	}
	// }()
	go func() {
		for d := range msgs {
			var envelope MessageEnvelope
			if err := json.Unmarshal(d.Body, &envelope); err != nil {
				log.Println("Failed to decode envelope:", err)
				continue
			}

			switch envelope.Type {
			case "PostLikeCreated":
				var event entities.PostLikeCreatedEvent
				if err := json.Unmarshal(envelope.Data, &event); err != nil {
					log.Println("Failed to decode PostLikeCreatedEvent:", err)
					continue
				}
				notificationUC.HandlePostLikeCreatedEvent(context.Background(), &event)

			case "CommentCreated":
				var event entities.CommentCreatedEvent
				if err := json.Unmarshal(envelope.Data, &event); err != nil {
					log.Println("Failed to decode CommentCreatedEvent:", err)
					continue
				}
				notificationUC.HandleCommentCreatedEvent(context.Background(), &event)

			case "EventCreated":
				var event entities.EventCreatedEvent
				if err := json.Unmarshal(envelope.Data, &event); err != nil {
					log.Println("Failed to decode EventCreatedEvent:", err)
					continue
				}
				notificationUC.HandleEventCreatedEvent(context.Background(), &event)

			case "UserFollowCreated":
				var event entities.UserFollowCreatedEvent
				if err := json.Unmarshal(envelope.Data, &event); err != nil {
					log.Println("Failed to decode UserFollowCreatedEvent:", err)
					continue
				}
				notificationUC.HandleUserFollowCreatedEvent(context.Background(), &event)

			default:
				log.Printf("Unknown event type: %s\n", envelope.Type)
			}
		}
	}()

}
