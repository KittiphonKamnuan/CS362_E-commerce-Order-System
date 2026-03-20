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

// GetNotifications handles GET /api/v1/notifications
func (h *NotificationHandler) GetNotifications(w http.ResponseWriter, r *http.Request) {
	customerID := r.Header.Get("X-Customer-ID")

	notifications, err := h.notificationService.GetNotifications(r.Context(), customerID)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(notifications))
}

// MarkAsRead handles PATCH /api/v1/notifications/{notificationId}/read
func (h *NotificationHandler) MarkAsRead(w http.ResponseWriter, r *http.Request) {
	notificationID := extractPathParam(r.URL.Path, "notifications")

	if err := h.notificationService.MarkAsRead(r.Context(), notificationID); err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(nil))
}
