package entity

import (
	"errors"
	"time"
)

// Payment records a payment transaction for an order
type Payment struct {
	PaymentID     string
	OrderID       string
	Amount        float64
	Method        PaymentMethod
	Status        PaymentStatus
	TransactionID string
	CreatedAt     time.Time
	PaidAt        *time.Time
	FailureReason string
}

// MarkCompleted transitions the payment to COMPLETED state
func (p *Payment) MarkCompleted(transactionID string, paidAt time.Time) error {
	if p.Status != PaymentStatusPending {
		return errors.New("only pending payments can be completed")
	}
	p.TransactionID = transactionID
	p.PaidAt = &paidAt
	p.Status = PaymentStatusCompleted
	return nil
}

// MarkFailed transitions the payment to FAILED state
func (p *Payment) MarkFailed(reason string) error {
	if p.Status != PaymentStatusPending {
		return errors.New("only pending payments can be marked failed")
	}
	p.FailureReason = reason
	p.Status = PaymentStatusFailed
	return nil
}

// Verify returns true if the payment is completed with a valid transaction ID
func (p *Payment) Verify() bool {
	return p.Status == PaymentStatusCompleted && p.TransactionID != ""
}

// IsValid returns an error if the payment is missing required fields
func (p *Payment) IsValid() error {
	if p.OrderID == "" {
		return errors.New("payment must reference an order")
	}
	if p.Amount <= 0 {
		return errors.New("payment amount must be positive")
	}
	return nil
}
