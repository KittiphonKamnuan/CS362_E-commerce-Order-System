package response

import "time"

// OrderItemResponse represents a single item in an order response
type OrderItemResponse struct {
	ProductID   string  `json:"productId"`
	ProductName string  `json:"productName"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
	Subtotal    float64 `json:"subtotal"`
}

// OrderResponse is the DTO returned for order queries
type OrderResponse struct {
	OrderID           string              `json:"orderId"`
	CustomerID        string              `json:"customerId"`
	Status            string              `json:"status"`
	TotalAmount       float64             `json:"totalAmount"`
	ShippingAddressID string              `json:"shippingAddressId"`
	Items             []OrderItemResponse `json:"items"`
	CreatedAt         time.Time           `json:"createdAt"`
	UpdatedAt         time.Time           `json:"updatedAt"`
}

// OrderListResponse wraps a paginated list of orders
type OrderListResponse struct {
	Orders   []*OrderResponse `json:"orders"`
	Total    int              `json:"total"`
	Page     int              `json:"page"`
	PageSize int              `json:"pageSize"`
}
