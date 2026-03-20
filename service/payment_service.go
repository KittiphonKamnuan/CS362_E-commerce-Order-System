package service

import (
	"context"

	"github.com/cs362/ecommerce/entity"
	"github.com/cs362/ecommerce/dto/request"
)

// PaymentService defines business logic for payment processing
type PaymentService interface {
	// InitiatePayment creates a new payment record for an order
	InitiatePayment(ctx context.Context, orderID string, req *request.InitiatePaymentRequest) (*entity.Payment, error)

	// ProcessPayment executes the payment through the payment gateway
	ProcessPayment(ctx context.Context, paymentID string) (*entity.Payment, error)

	// Refund initiates a refund for a completed payment
	Refund(ctx context.Context, paymentID string, reason string) (*entity.Payment, error)

	// GetPaymentByID retrieves a payment by its ID
	GetPaymentByID(ctx context.Context, paymentID string) (*entity.Payment, error)
}
