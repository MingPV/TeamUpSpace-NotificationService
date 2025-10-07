package entities

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	SendTo    uuid.UUID `json:"user_id"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
