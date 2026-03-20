package repository

import (
	"context"

	"github.com/cs362/ecommerce/entity"
)

// OrderRepository defines persistence operations for orders
type OrderRepository interface {
	// Create persists a new order and returns it with generated ID
	Create(ctx context.Context, order *entity.Order) (*entity.Order, error)

	// FindByID retrieves an order by its ID
	FindByID(ctx context.Context, orderID string) (*entity.Order, error)

	// FindByCustomerID returns paginated orders for a customer
	// Returns slice of orders, total count, and error
	FindByCustomerID(ctx context.Context, customerID string, offset, limit int) ([]*entity.Order, int, error)

	// UpdateStatus changes the status of an order
	UpdateStatus(ctx context.Context, orderID string, status entity.OrderStatus) error
}
