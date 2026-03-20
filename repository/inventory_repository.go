package repository

import (
	"context"

	"github.com/cs362/ecommerce/entity"
)

// InventoryRepository defines persistence operations for inventory
type InventoryRepository interface {
	// FindByProductID retrieves inventory record for a product
	FindByProductID(ctx context.Context, productID string) (*entity.Inventory, error)

	// ReserveStock atomically reserves stock using SELECT FOR UPDATE
	ReserveStock(ctx context.Context, productID string, qty int) error

	// ReleaseStock returns previously reserved stock back to available
	ReleaseStock(ctx context.Context, productID string, qty int) error

	// DeductStock permanently removes stock after order fulfillment
	DeductStock(ctx context.Context, productID string, qty int) error

	// AcquireLock creates a timed inventory lock for an order item
	AcquireLock(ctx context.Context, lock *entity.InventoryLock) (*entity.InventoryLock, error)
}
