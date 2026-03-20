package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cs362/ecommerce/dto/request"
	"github.com/cs362/ecommerce/dto/response"
	"github.com/cs362/ecommerce/service"
)

// CartHandler handles HTTP requests for cart endpoints
type CartHandler struct {
	cartService service.CartService
}

// NewCartHandler creates a new CartHandler
func NewCartHandler(cartService service.CartService) *CartHandler {
	return &CartHandler{cartService: cartService}
}

// GetCart handles GET /api/v1/cart
func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	customerID := r.Header.Get("X-Customer-ID")

	cart, err := h.cartService.GetCart(r.Context(), customerID)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(cart))
}

// AddItem handles POST /api/v1/cart/items
func (h *CartHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	customerID := r.Header.Get("X-Customer-ID")

	var req request.AddCartItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, response.NewError("BAD_REQUEST", err.Error()))
		return
	}

	cart, err := h.cartService.AddItem(r.Context(), customerID, &req)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(cart))
}

// UpdateItem handles PATCH /api/v1/cart/items/{productId}
func (h *CartHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	customerID := r.Header.Get("X-Customer-ID")
	productID := extractPathParam(r.URL.Path, "items")

	var req request.UpdateCartItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, response.NewError("BAD_REQUEST", err.Error()))
		return
	}

	cart, err := h.cartService.UpdateItemQty(r.Context(), customerID, productID, &req)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(cart))
}

// RemoveItem handles DELETE /api/v1/cart/items/{productId}
func (h *CartHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	customerID := r.Header.Get("X-Customer-ID")
	productID := extractPathParam(r.URL.Path, "items")

	cart, err := h.cartService.RemoveItem(r.Context(), customerID, productID)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(cart))
}
