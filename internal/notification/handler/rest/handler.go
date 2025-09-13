package rest

import (
	"strconv"

	"github.com/MingPV/NotificationService/internal/entities"
	"github.com/MingPV/NotificationService/internal/notification/dto"
	"github.com/MingPV/NotificationService/internal/notification/usecase"
	responses "github.com/MingPV/NotificationService/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpNotificationHandler struct {
	notificationUseCase usecase.NotificationUseCase
}

func NewHttpNotificationHandler(useCase usecase.NotificationUseCase) *HttpNotificationHandler {
	return &HttpNotificationHandler{notificationUseCase: useCase}
}

// CreateNotification godoc
// @Summary Create a new notification
// @Tags notifications
// @Accept json
// @Produce json
// @Param notification body entities.Notification true "Notification payload"
// @Success 201 {object} entities.Notification
// @Router /notifications [post]
func (h *HttpNotificationHandler) CreateNotification(c *fiber.Ctx) error {
	var req dto.CreateNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	notification := &entities.Notification{SendTo: req.SendTo, Type: req.Type, Message: req.Message}
	if err := h.notificationUseCase.CreateNotification(notification); err != nil {
		return responses.Error(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToNotificationResponse(notification))
}

// FindAllNotifications godoc
// @Summary Get all notifications
// @Tags notifications
// @Produce json
// @Success 200 {array} entities.Notification
// @Router /notifications [get]
func (h *HttpNotificationHandler) FindAllNotifications(c *fiber.Ctx) error {
	notifications, err := h.notificationUseCase.FindAllNotifications()
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToNotificationResponseList(notifications))
}

// FindNotificationByID godoc
// @Summary Get notification by ID
// @Tags notifications
// @Produce json
// @Param id path int true "Notification ID"
// @Success 200 {object} entities.Notification
// @Router /notifications/{id} [get]
func (h *HttpNotificationHandler) FindNotificationByID(c *fiber.Ctx) error {
	id := c.Params("id")
	notificationID, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	notification, err := h.notificationUseCase.FindNotificationByID(notificationID)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToNotificationResponse(notification))
}

// PatchNotification godoc
// @Summary Update an notification partially
// @Tags notifications
// @Accept json
// @Produce json
// @Param id path int true "Notification ID"
// @Param notification body entities.Notification true "Notification update payload"
// @Success 200 {object} entities.Notification
// @Router /notifications/{id} [patch]
func (h *HttpNotificationHandler) PatchNotification(c *fiber.Ctx) error {
	id := c.Params("id")
	notificationID, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	var req dto.CreateNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	notification := &entities.Notification{SendTo: req.SendTo, Type: req.Type, Message: req.Message}

	msg, err := validatePatchNotification(notification)
	if err != nil {
		return responses.ErrorWithMessage(c, err, msg)
	}

	updatedNotification, err := h.notificationUseCase.PatchNotification(notificationID, notification)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToNotificationResponse(updatedNotification))
}

// DeleteNotification godoc
// @Summary Delete an notification by ID
// @Tags notifications
// @Produce json
// @Param id path int true "Notification ID"
// @Success 200 {object} response.MessageResponse
// @Router /notifications/{id} [delete]
func (h *HttpNotificationHandler) DeleteNotification(c *fiber.Ctx) error {
	id := c.Params("id")
	notificationID, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	if err := h.notificationUseCase.DeleteNotification(notificationID); err != nil {
		return responses.Error(c, err)
	}

	return responses.Message(c, fiber.StatusOK, "notification deleted")
}

func validatePatchNotification(notification *entities.Notification) (string, error) {

	return "", nil
}
