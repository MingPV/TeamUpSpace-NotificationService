package usecase

import "github.com/MingPV/NotificationService/internal/entities"

type NotificationUseCase interface {
	FindAllNotifications() ([]*entities.Notification, error)
	CreateNotification(notification *entities.Notification) error
	PatchNotification(id int, notification *entities.Notification) (*entities.Notification, error)
	DeleteNotification(id int) error
	FindNotificationByID(id int) (*entities.Notification, error)
}