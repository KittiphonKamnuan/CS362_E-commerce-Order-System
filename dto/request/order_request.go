package request

// PlaceOrderRequest is the payload for POST /api/v1/orders
type PlaceOrderRequest struct {
	CartID            string `json:"cartId"`
	PaymentMethod     string `json:"paymentMethod"`
	ShippingAddressID string `json:"shippingAddressId"`
}
