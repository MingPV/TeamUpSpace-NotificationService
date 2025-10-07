package dto

import "github.com/google/uuid"

type CreateNotificationRequest struct {
	SendTo  uuid.UUID `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Type    string    `json:"type" validate:"required"`
	Message string    `json:"message" validate:"required"`
	IsRead  bool      `json:"is_read"`
}
