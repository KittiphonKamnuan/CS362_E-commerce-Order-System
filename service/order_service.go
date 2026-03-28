package service

import (
	"context"
	"time"
)

//
// ==========================
// Request DTO
// ==========================
//

// PlaceOrderRequest is the payload for creating an order
type PlaceOrderRequest struct {
	CartID            string `json:"cartId"`
	PaymentMethod     string `json:"paymentMethod"`
	ShippingAddressID string `json:"shippingAddressId"`
}

//
// ==========================
// Response DTO
// ==========================
//

// OrderItemResponse represents a single item in an order
type OrderItemResponse struct {
	ProductID   string  `json:"productId"`
	ProductName string  `json:"productName"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
	Subtotal    float64 `json:"subtotal"`
}

// OrderResponse represents an order returned to the client
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

// OrderListResponse represents paginated orders
type OrderListResponse struct {
	Orders   []*OrderResponse `json:"orders"`
	Total    int              `json:"total"`
	Page     int              `json:"page"`
	PageSize int              `json:"pageSize"`
}

//
// ==========================
// Service Interface
// ==========================
//

// OrderService defines business logic for order management
type OrderService interface {

	// PlaceOrder processes a flash sale order
	PlaceOrder(ctx context.Context, customerID string, req *PlaceOrderRequest) (*OrderResponse, error)

	// GetOrderByID retrieves an order by ID
	GetOrderByID(ctx context.Context, customerID, orderID string) (*OrderResponse, error)

	// GetOrderHistory retrieves paginated order history
	GetOrderHistory(ctx context.Context, customerID string, page, pageSize int) (*OrderListResponse, error)

	// CancelOrder cancels an order
	CancelOrder(ctx context.Context, customerID, orderID string) (*OrderResponse, error)
}
