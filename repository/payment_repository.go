package repository

import (
	"context"

	"github.com/cs362/ecommerce/entity"
)

// PaymentRepository defines persistence operations for payments
type PaymentRepository interface {
	// Create persists a new payment record
	Create(ctx context.Context, payment *entity.Payment) (*entity.Payment, error)

	// FindByID retrieves a payment by its ID
	FindByID(ctx context.Context, paymentID string) (*entity.Payment, error)

	// FindByOrderID retrieves all payments associated with an order
	FindByOrderID(ctx context.Context, orderID string) ([]*entity.Payment, error)

	// UpdateStatus updates the status and relevant fields of a payment
	Update(ctx context.Context, payment *entity.Payment) error
}
