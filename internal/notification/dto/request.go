package dto

type CreateNotificationRequest struct {
	SendTo  string `json:"send_to" validate:"required,email"`
	Type    string `json:"type" validate:"required"`
	Message string `json:"message" validate:"required"`
}