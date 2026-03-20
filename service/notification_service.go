package service

import (
	"context"

	"github.com/cs362/ecommerce/entity"
)

// NotificationService defines business logic for sending notifications
type NotificationService interface {
	// Send delivers a notification to a customer
	Send(ctx context.Context, customerID string, notifType entity.NotificationType, message string) error

	// SendBulk delivers the same notification to multiple customers
	SendBulk(ctx context.Context, customerIDs []string, notifType entity.NotificationType, message string) error

	// GetNotifications retrieves all notifications for a customer
	GetNotifications(ctx context.Context, customerID string) ([]*entity.Notification, error)

	// MarkAsRead marks a notification as read
	MarkAsRead(ctx context.Context, notificationID string) error
}
