package entities

import "time"

type Notification struct {
    ID    uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	SendTo    string    `json:"send_to"`
    Type      string    `json:"type"`
    Message   string    `json:"message"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}