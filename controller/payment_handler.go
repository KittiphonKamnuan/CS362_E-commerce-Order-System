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

// InitiatePayment godoc
//
//	@Summary		Initiate a payment
//	@Description	Creates a new payment record for an existing order and processes it through the payment gateway
//	@Tags			Payments
//	@Accept			json
//	@Produce		json
//	@Param			body	body		request.InitiatePaymentRequest	true	"Payment request payload"
//	@Success		201		{object}	response.SuccessResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		404		{object}	response.ErrorResponse	"Order not found"
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/payments [post]
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

// Refund godoc
//
//	@Summary		Refund a payment
//	@Description	Initiates a refund for a completed payment. Only COMPLETED payments can be refunded.
//	@Tags			Payments
//	@Accept			json
//	@Produce		json
//	@Param			paymentId	path		string					true	"Payment ID"
//	@Param			body		body		request.RefundRequest	true	"Refund reason"
//	@Success		200			{object}	response.SuccessResponse
//	@Failure		400			{object}	response.ErrorResponse
//	@Failure		404			{object}	response.ErrorResponse
//	@Failure		409			{object}	response.ErrorResponse	"Payment not eligible for refund"
//	@Failure		500			{object}	response.ErrorResponse
//	@Router			/payments/{paymentId}/refund [post]
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
