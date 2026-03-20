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

// GetCart godoc
//
//	@Summary		Get customer's cart
//	@Description	Retrieves the active shopping cart for the authenticated customer
//	@Tags			Cart
//	@Produce		json
//	@Security		CustomerID
//	@Param			X-Customer-ID	header		string	true	"Customer ID"
//	@Success		200				{object}	response.SuccessResponse
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/cart [get]
func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	customerID := r.Header.Get("X-Customer-ID")

	cart, err := h.cartService.GetCart(r.Context(), customerID)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(cart))
}

// AddItem godoc
//
//	@Summary		Add item to cart
//	@Description	Adds a product to the customer's cart. Increments quantity if product already exists.
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Security		CustomerID
//	@Param			X-Customer-ID	header		string						true	"Customer ID"
//	@Param			body			body		request.AddCartItemRequest	true	"Product and quantity"
//	@Success		200				{object}	response.SuccessResponse
//	@Failure		400				{object}	response.ErrorResponse
//	@Failure		404				{object}	response.ErrorResponse	"Product not found"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/cart/items [post]
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

// UpdateItem godoc
//
//	@Summary		Update item quantity in cart
//	@Description	Updates the quantity of a specific product in the cart. Set quantity to 0 to remove.
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Security		CustomerID
//	@Param			X-Customer-ID	header		string							true	"Customer ID"
//	@Param			productId		path		string							true	"Product ID"
//	@Param			body			body		request.UpdateCartItemRequest	true	"New quantity"
//	@Success		200				{object}	response.SuccessResponse
//	@Failure		400				{object}	response.ErrorResponse
//	@Failure		404				{object}	response.ErrorResponse
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/cart/items/{productId} [patch]
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

// RemoveItem godoc
//
//	@Summary		Remove item from cart
//	@Description	Removes a specific product from the customer's cart
//	@Tags			Cart
//	@Produce		json
//	@Security		CustomerID
//	@Param			X-Customer-ID	header		string	true	"Customer ID"
//	@Param			productId		path		string	true	"Product ID"
//	@Success		200				{object}	response.SuccessResponse
//	@Failure		404				{object}	response.ErrorResponse
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/cart/items/{productId} [delete]
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
