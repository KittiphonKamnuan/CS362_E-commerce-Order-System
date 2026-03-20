package entity

import "time"

// Notification represents a message sent to a customer
type Notification struct {
	NotificationID string
	CustomerID     string
	Type           NotificationType
	Message        string
	SentAt         time.Time
	IsRead         bool
}
