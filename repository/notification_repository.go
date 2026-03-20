package repository

import (
	"context"

	"github.com/cs362/ecommerce/entity"
)

// NotificationRepository defines persistence operations for notifications
type NotificationRepository interface {
	// Create persists a new notification
	Create(ctx context.Context, notification *entity.Notification) (*entity.Notification, error)

	// FindByCustomerID returns all notifications for a customer
	FindByCustomerID(ctx context.Context, customerID string) ([]*entity.Notification, error)

	// MarkAsRead marks a notification as read
	MarkAsRead(ctx context.Context, notificationID string) error
}
