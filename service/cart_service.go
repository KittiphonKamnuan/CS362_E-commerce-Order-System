package service

import (
	"context"

	"github.com/cs362/ecommerce/entity"
	"github.com/cs362/ecommerce/dto/request"
)

// CartService defines business logic for shopping cart management
type CartService interface {
	// GetCart retrieves the active cart for a customer
	GetCart(ctx context.Context, customerID string) (*entity.Cart, error)

	// AddItem adds a product to the customer's cart
	AddItem(ctx context.Context, customerID string, req *request.AddCartItemRequest) (*entity.Cart, error)

	// UpdateItemQty changes the quantity of an existing cart item
	UpdateItemQty(ctx context.Context, customerID, productID string, req *request.UpdateCartItemRequest) (*entity.Cart, error)

	// RemoveItem removes a product from the customer's cart
	RemoveItem(ctx context.Context, customerID, productID string) (*entity.Cart, error)

	// ClearCart empties the customer's cart after order placement
	ClearCart(ctx context.Context, customerID string) error
}
