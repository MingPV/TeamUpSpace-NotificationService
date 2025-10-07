package dto

import (
	"time"

	"github.com/google/uuid"
)

type NotificationResponse struct {
	ID        uint      `json:"id"`
	SendTo    uuid.UUID `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
