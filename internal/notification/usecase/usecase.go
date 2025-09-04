package usecase

import (
	"github.com/MingPV/NotificationService/internal/entities"
	"github.com/MingPV/NotificationService/internal/notification/repository"
)

// NotificationService
type NotificationService struct {
	repo repository.NotificationRepository
}

// Init NotificationService function
func NewNotificationService(repo repository.NotificationRepository) NotificationUseCase {
	return &NotificationService{repo: repo}
}

// NotificationService Methods - 1 create
func (s *NotificationService) CreateNotification(notification *entities.Notification) error {
	if err := s.repo.Save(notification); err != nil {
		return err
	}
	return nil
}

// NotificationService Methods - 2 find all
func (s *NotificationService) FindAllNotifications() ([]*entities.Notification, error) {
	notifications, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

// NotificationService Methods - 3 find by id
func (s *NotificationService) FindNotificationByID(id int) (*entities.Notification, error) {

	notification, err := s.repo.FindByID(id)
	if err != nil {
		return &entities.Notification{}, err
	}
	return notification, nil
}

// NotificationService Methods - 4 patch
func (s *NotificationService) PatchNotification(id int, notification *entities.Notification) (*entities.Notification, error) {

	if err := s.repo.Patch(id, notification); err != nil {
		return nil, err
	}
	updatedNotification, _ := s.repo.FindByID(id)

	return updatedNotification, nil
}

// NotificationService Methods - 5 delete
func (s *NotificationService) DeleteNotification(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
