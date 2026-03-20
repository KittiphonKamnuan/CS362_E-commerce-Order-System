package request

// AddCartItemRequest is the payload for POST /api/v1/cart/items
type AddCartItemRequest struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

// UpdateCartItemRequest is the payload for PATCH /api/v1/cart/items/{productId}
type UpdateCartItemRequest struct {
	Quantity int `json:"quantity"`
}
