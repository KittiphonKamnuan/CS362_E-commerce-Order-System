package service

import (
	"context"

	"github.com/cs362/ecommerce/dto/request"
	"github.com/cs362/ecommerce/dto/response"
)

// OrderService defines business logic for order management
type OrderService interface {
	// PlaceOrder processes a flash sale order for a customer
	PlaceOrder(ctx context.Context, customerID string, req *request.PlaceOrderRequest) (*response.OrderResponse, error)

	// GetOrderByID retrieves a specific order owned by the customer
	GetOrderByID(ctx context.Context, customerID, orderID string) (*response.OrderResponse, error)

	// GetOrderHistory returns paginated order history for a customer
	GetOrderHistory(ctx context.Context, customerID string, page, pageSize int) (*response.OrderListResponse, error)

	// CancelOrder cancels an order if it is still in a cancellable state
	CancelOrder(ctx context.Context, customerID, orderID string) (*response.OrderResponse, error)
}
