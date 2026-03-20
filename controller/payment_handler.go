package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cs362/ecommerce/dto/request"
	"github.com/cs362/ecommerce/dto/response"
	"github.com/cs362/ecommerce/service"
)

// PaymentHandler handles HTTP requests for payment endpoints
type PaymentHandler struct {
	paymentService service.PaymentService
}

// NewPaymentHandler creates a new PaymentHandler
func NewPaymentHandler(paymentService service.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

// InitiatePayment handles POST /api/v1/payments
func (h *PaymentHandler) InitiatePayment(w http.ResponseWriter, r *http.Request) {
	var req request.InitiatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, response.NewError("BAD_REQUEST", err.Error()))
		return
	}

	payment, err := h.paymentService.InitiatePayment(r.Context(), req.OrderID, &req)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, response.NewSuccess(payment))
}

// Refund handles POST /api/v1/payments/{paymentId}/refund
func (h *PaymentHandler) Refund(w http.ResponseWriter, r *http.Request) {
	paymentID := extractPathParam(r.URL.Path, "payments")

	var req request.RefundRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, response.NewError("BAD_REQUEST", err.Error()))
		return
	}

	payment, err := h.paymentService.Refund(r.Context(), paymentID, req.Reason)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(payment))
}
