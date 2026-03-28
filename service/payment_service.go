package service

import (
	"context"
	"errors"
	"time"
)

//////////////////////
// ENUM DEFINITIONS //
//////////////////////

type PaymentMethod string

const (
	PaymentMethodCreditCard PaymentMethod = "CREDIT_CARD"
	PaymentMethodPromptPay  PaymentMethod = "PROMPTPAY"
	PaymentMethodCOD        PaymentMethod = "COD"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusCompleted PaymentStatus = "COMPLETED"
	PaymentStatusFailed    PaymentStatus = "FAILED"
	PaymentStatusRefunded  PaymentStatus = "REFUNDED"
)

//////////////////
// ENTITY MODEL //
//////////////////

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

func (p *Payment) MarkCompleted(transactionID string, paidAt time.Time) error {
	if p.Status != PaymentStatusPending {
		return errors.New("only pending payments can be completed")
	}

	p.TransactionID = transactionID
	p.PaidAt = &paidAt
	p.Status = PaymentStatusCompleted

	return nil
}

func (p *Payment) MarkFailed(reason string) error {
	if p.Status != PaymentStatusPending {
		return errors.New("only pending payments can be marked failed")
	}

	p.FailureReason = reason
	p.Status = PaymentStatusFailed

	return nil
}

func (p *Payment) Verify() bool {
	return p.Status == PaymentStatusCompleted && p.TransactionID != ""
}

func (p *Payment) IsValid() error {
	if p.OrderID == "" {
		return errors.New("payment must reference an order")
	}

	if p.Amount <= 0 {
		return errors.New("payment amount must be positive")
	}

	return nil
}

/////////////////////
// REQUEST OBJECTS //
/////////////////////

type InitiatePaymentRequest struct {
	OrderID       string `json:"orderId"`
	PaymentMethod string `json:"paymentMethod"`
}

type RefundRequest struct {
	Reason string `json:"reason"`
}

/////////////////////
// SERVICE LAYER   //
/////////////////////

type PaymentService interface {

	// create payment record
	InitiatePayment(
		ctx context.Context,
		orderID string,
		req *InitiatePaymentRequest,
	) (*Payment, error)

	// send payment to gateway
	ProcessPayment(
		ctx context.Context,
		paymentID string,
	) (*Payment, error)

	// refund completed payment
	Refund(
		ctx context.Context,
		paymentID string,
		reason string,
	) (*Payment, error)

	// get payment by id
	GetPaymentByID(
		ctx context.Context,
		paymentID string,
	) (*Payment, error)
}