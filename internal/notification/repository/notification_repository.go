package repository

import "github.com/MingPV/NotificationService/internal/entities"

type NotificationRepository interface {
	Save(notification *entities.Notification) error
	FindAll() ([]*entities.Notification, error)
	FindByID(id int) (*entities.Notification, error)
	FindByUserID(userID string) ([]*entities.Notification, error)
	Patch(id int, notification *entities.Notification) error
	Delete(id int) error
}
