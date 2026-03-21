package handler

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

// PlaceOrder godoc
//
//	@Summary		Place a new order (Flash Sale)
//	@Description	Creates an order from the customer's cart. Enforces rate limiting for flash sale. Returns 429 if rate limit exceeded, 409 if stock is insufficient.
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Security		CustomerID
//	@Param			X-Customer-ID	header		string					true	"Customer ID"
//	@Param			body			body		request.PlaceOrderRequest	true	"Order request payload"
//	@Success		201				{object}	response.SuccessResponse{data=response.OrderResponse}
//	@Failure		400				{object}	response.ErrorResponse
//	@Failure		409				{object}	response.ErrorResponse	"INSUFFICIENT_STOCK"
//	@Failure		429				{object}	response.ErrorResponse	"RATE_LIMIT_EXCEEDED"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/orders [post]
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

// GetOrder godoc
//
//	@Summary		Get order by ID
//	@Description	Retrieves a specific order belonging to the authenticated customer
//	@Tags			Orders
//	@Produce		json
//	@Security		CustomerID
//	@Param			X-Customer-ID	header		string	true	"Customer ID"
//	@Param			orderId			path		string	true	"Order ID"
//	@Success		200				{object}	response.SuccessResponse{data=response.OrderResponse}
//	@Failure		404				{object}	response.ErrorResponse
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/orders/{orderId} [get]
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

// ListOrders godoc
//
//	@Summary		List orders for a customer
//	@Description	Returns paginated order history for the authenticated customer
//	@Tags			Orders
//	@Produce		json
//	@Security		CustomerID
//	@Param			X-Customer-ID	header		string	true	"Customer ID"
//	@Param			page			query		int		false	"Page number"		default(1)
//	@Param			pageSize		query		int		false	"Page size"			default(10)
//	@Success		200				{object}	response.SuccessResponse{data=response.OrderListResponse}
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/orders [get]
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

// CancelOrder godoc
//
//	@Summary		Cancel an order
//	@Description	Cancels an order if it is in PENDING or CONFIRMED state
//	@Tags			Orders
//	@Produce		json
//	@Security		CustomerID
//	@Param			X-Customer-ID	header		string	true	"Customer ID"
//	@Param			orderId			path		string	true	"Order ID"
//	@Success		200				{object}	response.SuccessResponse{data=response.OrderResponse}
//	@Failure		404				{object}	response.ErrorResponse
//	@Failure		409				{object}	response.ErrorResponse	"Order cannot be cancelled"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/orders/{orderId}/cancel [patch]
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
