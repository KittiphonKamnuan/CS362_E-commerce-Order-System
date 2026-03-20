package repository

import (
	"context"

	"github.com/cs362/ecommerce/entity"
)

// CartRepository defines persistence operations for shopping carts
type CartRepository interface {
	// FindByCustomerID retrieves the active cart for a customer
	FindByCustomerID(ctx context.Context, customerID string) (*entity.Cart, error)

	// FindByID retrieves a cart by its ID
	FindByID(ctx context.Context, cartID string) (*entity.Cart, error)

	// Save persists cart state (upsert)
	Save(ctx context.Context, cart *entity.Cart) (*entity.Cart, error)

	// Clear removes all items from a cart
	Clear(ctx context.Context, cartID string) error
}
