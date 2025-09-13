package grpc

import (
    "context"

    "github.com/MingPV/NotificationService/internal/entities"
    "github.com/MingPV/NotificationService/internal/notification/usecase"
    "github.com/MingPV/NotificationService/pkg/apperror"
    notificationpb "github.com/MingPV/NotificationService/proto/notification"
    "github.com/google/uuid"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcNotificationHandler struct {
	notificationUseCase usecase.NotificationUseCase
	notificationpb.UnimplementedNotificationServiceServer
}

func NewGrpcNotificationHandler(uc usecase.NotificationUseCase) *GrpcNotificationHandler {
	return &GrpcNotificationHandler{notificationUseCase: uc}
}

func (h *GrpcNotificationHandler) CreateNotification(ctx context.Context, req *notificationpb.CreateNotificationRequest) (*notificationpb.CreateNotificationResponse, error) {
	sendToUUID, _ := uuid.Parse(req.SendTo) // assume validation done elsewhere
	notif := &entities.Notification{
		SendTo:  sendToUUID,
		Type:    req.Type,
		Message: req.Message,
	}

	if err := h.notificationUseCase.CreateNotification(notif); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &notificationpb.CreateNotificationResponse{Notification: toProtoNotification(notif)}, nil
}

func (h *GrpcNotificationHandler) FindNotificationByID(ctx context.Context, req *notificationpb.FindNotificationByIDRequest) (*notificationpb.FindNotificationByIDResponse, error) {
	notif, err := h.notificationUseCase.FindNotificationByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &notificationpb.FindNotificationByIDResponse{Notification: toProtoNotification(notif)}, nil
}

func (h *GrpcNotificationHandler) FindAllNotifications(ctx context.Context, req *notificationpb.FindAllNotificationsRequest) (*notificationpb.FindAllNotificationsResponse, error) {
	notifs, err := h.notificationUseCase.FindAllNotifications()
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoNotifs []*notificationpb.Notification
	for _, n := range notifs {
		protoNotifs = append(protoNotifs, toProtoNotification(n))
	}

	return &notificationpb.FindAllNotificationsResponse{Notifications: protoNotifs}, nil
}

func (h *GrpcNotificationHandler) PatchNotification(ctx context.Context, req *notificationpb.PatchNotificationRequest) (*notificationpb.PatchNotificationResponse, error) {
	sendToUUID, _ := uuid.Parse(req.SendTo)
	notif := &entities.Notification{
		SendTo:  sendToUUID,
		Type:    req.Type,
		Message: req.Message,
	}

	updatedNotif, err := h.notificationUseCase.PatchNotification(int(req.Id), notif)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &notificationpb.PatchNotificationResponse{Notification: toProtoNotification(updatedNotif)}, nil
}

func (h *GrpcNotificationHandler) DeleteNotification(ctx context.Context, req *notificationpb.DeleteNotificationRequest) (*notificationpb.DeleteNotificationResponse, error) {
	if err := h.notificationUseCase.DeleteNotification(int(req.Id)); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &notificationpb.DeleteNotificationResponse{Message: "notification deleted"}, nil
}

// helper function convert entities.Notification to notificationpb.Notification
func toProtoNotification(n *entities.Notification) *notificationpb.Notification {
	return &notificationpb.Notification{
		Id:        int32(n.ID),
		SendTo:    n.SendTo.String(),
		Type:      n.Type,
		Message:   n.Message,
		CreatedAt: timestamppb.New(n.CreatedAt),
		UpdatedAt: timestamppb.New(n.UpdatedAt),
	}
}