package usecase

import (
	"context"
	"log"

	"github.com/MingPV/NotificationService/internal/entities"
	"github.com/MingPV/NotificationService/internal/notification/repository"
	"github.com/google/uuid"
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

func (s *NotificationService) FindNotificationsByUserID(userID string) ([]*entities.Notification, error) {
	notifications, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (s *NotificationService) MarkAsReadByUserID(userID string) error {
	notifications, err := s.repo.FindByUserID(userID)
	if err != nil {
		return err
	}

	for _, notification := range notifications {
		notification.IsRead = true
		if err := s.repo.Patch(int(notification.ID), notification); err != nil {
			log.Printf("Failed to mark notification ID %d as read: %v", notification.ID, err)
		}
	}

	return nil
}

// -----------------------------
// RabbitMQ consumer
// -----------------------------
func (s *NotificationService) HandleEventCreatedEvent(ctx context.Context, event *entities.EventCreatedEvent) {
	notification := &entities.Notification{
		// uuid that means all users
		SendTo:  uuid.Nil,
		Type:    "EventCreated",
		Message: "A new event has been created: " + event.EventName,
	}

	if err := s.repo.Save(notification); err != nil {
		log.Println("Failed to save notification from event:", err)
	} else {
		log.Println("Notification created from EventCreatedEvent:", notification)
	}
}

func (s *NotificationService) HandlePostLikeCreatedEvent(ctx context.Context, pl *entities.PostLikeCreatedEvent) {
	notification := &entities.Notification{
		SendTo:  pl.PostOwnerId,
		Type:    "PostLikeCreated",
		Message: "Your post has a new like",
	}

	if err := s.repo.Save(notification); err != nil {
		log.Println("Failed to save notification from PostLikeCreatedEvent:", err)
	} else {
		log.Println("Notification created from PostLikeCreatedEvent:", notification)
	}
}

func (s *NotificationService) HandleCommentCreatedEvent(ctx context.Context, cm *entities.CommentCreatedEvent) {
	notification := &entities.Notification{
		SendTo:  cm.PostOwnerId,
		Type:    "CommentCreated",
		Message: "Your post has a new comment",
	}

	if err := s.repo.Save(notification); err != nil {
		log.Println("Failed to save notification from CommentCreatedEvent:", err)
	} else {
		log.Println("Notification created from CommentCreatedEvent:", notification)
	}
}

func (s *NotificationService) HandleUserFollowCreatedEvent(ctx context.Context, uf *entities.UserFollowCreatedEvent) {
	notification := &entities.Notification{
		SendTo:  uf.FollowTo,
		Type:    "UserFollowCreated",
		Message: "You have a new follower",
	}
	if err := s.repo.Save(notification); err != nil {
		log.Println("Failed to save notification from UserFollowCreatedEvent:", err)
	} else {
		log.Println("Notification created from UserFollowCreatedEvent:", notification)
	}

}
