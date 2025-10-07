package usecase

import (
	"context"

	"github.com/MingPV/NotificationService/internal/entities"
)

type NotificationUseCase interface {
	FindAllNotifications() ([]*entities.Notification, error)
	CreateNotification(notification *entities.Notification) error
	PatchNotification(id int, notification *entities.Notification) (*entities.Notification, error)
	DeleteNotification(id int) error
	FindNotificationByID(id int) (*entities.Notification, error)
	FindNotificationsByUserID(userID string) ([]*entities.Notification, error)
	MarkAsReadByUserID(userID string) error

	// MessageBroker
	HandleEventCreatedEvent(ctx context.Context, event *entities.EventCreatedEvent)
	HandlePostLikeCreatedEvent(ctx context.Context, event *entities.PostLikeCreatedEvent)
	HandleCommentCreatedEvent(ctx context.Context, event *entities.CommentCreatedEvent)
	HandleUserFollowCreatedEvent(ctx context.Context, event *entities.UserFollowCreatedEvent)
}
