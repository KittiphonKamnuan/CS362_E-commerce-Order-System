package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/cs362/ecommerce/dto/request"
	"github.com/cs362/ecommerce/dto/response"
	"github.com/cs362/ecommerce/service"
)

// OrderHandler handles HTTP requests for order endpoints
type OrderHandler struct {
	orderService service.OrderService
}

// NewOrderHandler creates a new OrderHandler
func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

// PlaceOrder handles POST /api/v1/orders
func (h *OrderHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.Header.Get("X-Customer-ID")

	var req request.PlaceOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, response.NewError("BAD_REQUEST", err.Error()))
		return
	}

	order, err := h.orderService.PlaceOrder(r.Context(), customerID, &req)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, response.NewSuccess(order))
}

// GetOrder handles GET /api/v1/orders/{orderId}
func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.Header.Get("X-Customer-ID")
	orderID := extractPathParam(r.URL.Path, "orders")

	order, err := h.orderService.GetOrderByID(r.Context(), customerID, orderID)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(order))
}

// ListOrders handles GET /api/v1/orders?page=1&pageSize=10
func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	customerID := r.Header.Get("X-Customer-ID")
	page := queryInt(r, "page", 1)
	pageSize := queryInt(r, "pageSize", 10)

	orders, err := h.orderService.GetOrderHistory(r.Context(), customerID, page, pageSize)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(orders))
}

// CancelOrder handles PATCH /api/v1/orders/{orderId}/cancel
func (h *OrderHandler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.Header.Get("X-Customer-ID")
	orderID := extractPathParam(r.URL.Path, "orders")

	order, err := h.orderService.CancelOrder(r.Context(), customerID, orderID)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(order))
}

// helper: write JSON response
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// helper: extract path segment after the given resource name
func extractPathParam(path, resource string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	for i, p := range parts {
		if p == resource && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}

// helper: parse integer query parameter with a default value
func queryInt(r *http.Request, key string, defaultVal int) int {
	v := r.URL.Query().Get(key)
	if v == "" {
		return defaultVal
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return defaultVal
	}
	return n
}

// handleServiceError maps known error strings to HTTP status codes
func handleServiceError(w http.ResponseWriter, err error) {
	msg := err.Error()
	switch {
	case strings.Contains(msg, "RATE_LIMIT_EXCEEDED"):
		writeJSON(w, http.StatusTooManyRequests, &response.ErrorResponse{
			Success: false,
			Error:   response.ErrorDetail{Code: "RATE_LIMIT_EXCEEDED", RetryAfter: 60},
		})
	case strings.Contains(msg, "INSUFFICIENT_STOCK"):
		writeJSON(w, http.StatusConflict, response.NewError("INSUFFICIENT_STOCK", msg))
	case strings.Contains(msg, "not found"):
		writeJSON(w, http.StatusNotFound, response.NewError("NOT_FOUND", msg))
	default:
		writeJSON(w, http.StatusInternalServerError, response.NewError("INTERNAL_ERROR", msg))
	}
}
