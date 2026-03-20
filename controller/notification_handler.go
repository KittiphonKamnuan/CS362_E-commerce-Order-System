package controller

import (
	"net/http"

	"github.com/cs362/ecommerce/dto/response"
	"github.com/cs362/ecommerce/service"
)

// NotificationHandler handles HTTP requests for notification endpoints
type NotificationHandler struct {
	notificationService service.NotificationService
}

// NewNotificationHandler creates a new NotificationHandler
func NewNotificationHandler(notificationService service.NotificationService) *NotificationHandler {
	return &NotificationHandler{notificationService: notificationService}
}

// GetNotifications godoc
//
//	@Summary		Get notifications for a customer
//	@Description	Returns all notifications for the authenticated customer
//	@Tags			Notifications
//	@Produce		json
//	@Security		CustomerID
//	@Param			X-Customer-ID	header		string	true	"Customer ID"
//	@Success		200				{object}	response.SuccessResponse
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/notifications [get]
func (h *NotificationHandler) GetNotifications(w http.ResponseWriter, r *http.Request) {
	customerID := r.Header.Get("X-Customer-ID")

	notifications, err := h.notificationService.GetNotifications(r.Context(), customerID)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(notifications))
}

// MarkAsRead godoc
//
//	@Summary		Mark notification as read
//	@Description	Marks a specific notification as read for the customer
//	@Tags			Notifications
//	@Produce		json
//	@Security		CustomerID
//	@Param			X-Customer-ID		header		string	true	"Customer ID"
//	@Param			notificationId		path		string	true	"Notification ID"
//	@Success		200					{object}	response.SuccessResponse
//	@Failure		404					{object}	response.ErrorResponse
//	@Failure		500					{object}	response.ErrorResponse
//	@Router			/notifications/{notificationId}/read [patch]
func (h *NotificationHandler) MarkAsRead(w http.ResponseWriter, r *http.Request) {
	notificationID := extractPathParam(r.URL.Path, "notifications")

	if err := h.notificationService.MarkAsRead(r.Context(), notificationID); err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(nil))
}
