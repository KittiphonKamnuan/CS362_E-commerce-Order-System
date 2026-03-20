package service

import (
	"context"

	"github.com/cs362/ecommerce/entity"
)

// InventoryService defines business logic for stock management
type InventoryService interface {
	// ValidateAndReserve checks stock availability and reserves the requested quantity
	ValidateAndReserve(ctx context.Context, productID string, qty int) error

	// AcquireLock creates a TTL-based inventory lock (15 min) for an order item
	AcquireLock(ctx context.Context, orderID, productID, customerID string, qty int) (*entity.InventoryLock, error)

	// ReleaseLock releases an existing inventory lock
	ReleaseLock(ctx context.Context, lockID string) error

	// FulfillOrder permanently deducts stock after payment confirmation
	FulfillOrder(ctx context.Context, orderID string) error
}
