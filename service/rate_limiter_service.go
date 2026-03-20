package service

import "context"

// RateLimiterService enforces request rate limits for flash sale orders
type RateLimiterService interface {
	// IsLimitExceeded checks whether the customer has exceeded the allowed rate
	IsLimitExceeded(ctx context.Context, customerID string) (bool, error)

	// ValidateOrderRequest returns an error if the customer is rate-limited
	ValidateOrderRequest(ctx context.Context, customerID string) error

	// RecordRequest increments the request counter for a customer
	RecordRequest(ctx context.Context, customerID string) error

	// ResetCounter clears the request counter for a customer
	ResetCounter(ctx context.Context, customerID string) error
}
