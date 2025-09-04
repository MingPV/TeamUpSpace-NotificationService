package dto

import "github.com/MingPV/NotificationService/internal/entities"

func ToNotificationResponse(notification *entities.Notification) *NotificationResponse {
	return &NotificationResponse{
		ID:    		notification.ID,
		SendTo: 	notification.SendTo,
		Type: 		notification.Type,
		Message: 	notification.Message,
		CreatedAt: 	notification.CreatedAt,
		UpdatedAt: 	notification.UpdatedAt,
	}
}

func ToNotificationResponseList(notifications []*entities.Notification) []*NotificationResponse {
	result := make([]*NotificationResponse, 0, len(notifications))
	for _, o := range notifications {
		result = append(result, ToNotificationResponse(o))
	}
	return result
}
