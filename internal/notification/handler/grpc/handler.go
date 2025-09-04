package grpc

import (
	"context"

	"github.com/MingPV/NotificationService/internal/entities"
	"github.com/MingPV/NotificationService/internal/notification/usecase"
	"github.com/MingPV/NotificationService/pkg/apperror"
	notificationpb "github.com/MingPV/NotificationService/proto/notification"
	"google.golang.org/grpc/status"
)

type GrpcNotificationHandler struct {
	notificationUseCase usecase.NotificationUseCase
	notificationpb.UnimplementedNotificationServiceServer
}

func NewGrpcNotificationHandler(uc usecase.NotificationUseCase) *GrpcNotificationHandler {
	return &GrpcNotificationHandler{notificationUseCase: uc}
}

func (h *GrpcNotificationHandler) CreateNotification(ctx context.Context, req *notificationpb.CreateNotificationRequest) (*notificationpb.CreateNotificationResponse, error) {
	notification := &entities.Notification{SendTo:  req.SendTo,Type:    req.Type,Message: req.Message}
	if err := h.notificationUseCase.CreateNotification(notification); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &notificationpb.CreateNotificationResponse{Notification: toProtoNotification(notification)}, nil
}

func (h *GrpcNotificationHandler) FindNotificationByID(ctx context.Context, req *notificationpb.FindNotificationByIDRequest) (*notificationpb.FindNotificationByIDResponse, error) {
	notification, err := h.notificationUseCase.FindNotificationByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &notificationpb.FindNotificationByIDResponse{Notification: toProtoNotification(notification)}, nil
}

func (h *GrpcNotificationHandler) FindAllNotifications(ctx context.Context, req *notificationpb.FindAllNotificationsRequest) (*notificationpb.FindAllNotificationsResponse, error) {
	notifications, err := h.notificationUseCase.FindAllNotifications()
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoNotifications []*notificationpb.Notification
	for _, o := range notifications {
		protoNotifications = append(protoNotifications, toProtoNotification(o))
	}

	return &notificationpb.FindAllNotificationsResponse{Notifications: protoNotifications}, nil
}

func (h *GrpcNotificationHandler) PatchNotification(ctx context.Context, req *notificationpb.PatchNotificationRequest) (*notificationpb.PatchNotificationResponse, error) {
	notification := &entities.Notification{Total: float64(req.Total)}
	updatedNotification, err := h.notificationUseCase.PatchNotification(int(req.Id), notification)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &notificationpb.PatchNotificationResponse{Notification: toProtoNotification(updatedNotification)}, nil
}

func (h *GrpcNotificationHandler) DeleteNotification(ctx context.Context, req *notificationpb.DeleteNotificationRequest) (*notificationpb.DeleteNotificationResponse, error) {
	if err := h.notificationUseCase.DeleteNotification(int(req.Id)); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &notificationpb.DeleteNotificationResponse{Message: "notification deleted"}, nil
}

// helper function convert entities.Notification to notificationpb.Notification
func toProtoNotification(o *entities.Notification) *notificationpb.Notification {
	return &notificationpb.Notification{
		Id:      int32(o.ID),
        SendTo:  o.SendTo,
        Type:    o.Type,
        Message: o.Message,
	}
}
