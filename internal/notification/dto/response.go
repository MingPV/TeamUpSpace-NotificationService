package dto

import "time"

type NotificationResponse struct {
	ID    uint    `json:"id"`
	SendTo    string    `json:"send_to"`
    Type      string    `json:"type"`
    Message   string    `json:"message"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
