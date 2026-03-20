package request

// InitiatePaymentRequest is the payload for POST /api/v1/payments
type InitiatePaymentRequest struct {
	OrderID       string `json:"orderId"`
	PaymentMethod string `json:"paymentMethod"`
}

// RefundRequest is the payload for POST /api/v1/payments/{paymentId}/refund
type RefundRequest struct {
	Reason string `json:"reason"`
}
