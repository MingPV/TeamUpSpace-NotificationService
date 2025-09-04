package repository

import (
	"github.com/MingPV/NotificationService/internal/entities"
	"gorm.io/gorm"
)

type GormNotificationRepository struct {
	db *gorm.DB
}

func NewGormNotificationRepository(db *gorm.DB) NotificationRepository {
	return &GormNotificationRepository{db: db}
}

func (r *GormNotificationRepository) Save(notification *entities.Notification) error {
	return r.db.Create(&notification).Error
}

func (r *GormNotificationRepository) FindAll() ([]*entities.Notification, error) {
	var notificationValues []entities.Notification
	if err := r.db.Find(&notificationValues).Error; err != nil {
		return nil, err
	}

	notifications := make([]*entities.Notification, len(notificationValues))
	for i := range notificationValues {
		notifications[i] = &notificationValues[i]
	}
	return notifications, nil
}

func (r *GormNotificationRepository) FindByID(id int) (*entities.Notification, error) {
	var notification entities.Notification
	if err := r.db.First(&notification, id).Error; err != nil {
		return &entities.Notification{}, err
	}
	return &notification, nil
}

func (r *GormNotificationRepository) Patch(id int, notification *entities.Notification) error {
	result := r.db.Model(&entities.Notification{}).Where("id = ?", id).Updates(notification)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormNotificationRepository) Delete(id int) error {
	result := r.db.Delete(&entities.Notification{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
