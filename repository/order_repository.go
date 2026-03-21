package repository

import "context"

// OrderRepository defines database operations for Order
type OrderRepository interface {
	Create(ctx context.Context, order *OrderModel) (*OrderModel, error)
	FindByID(ctx context.Context, orderID string) (*OrderModel, error)
	FindByCustomerID(ctx context.Context, customerID string, offset int, limit int) ([]*OrderModel, int, error)
	UpdateStatus(ctx context.Context, orderID string, status string) error
}

// InventoryRepository defines database operations for Inventory
type InventoryRepository interface {
	FindByProductID(ctx context.Context, productID string) (*InventoryModel, error)
	ReserveStock(ctx context.Context, productID string, qty int) error
	ReleaseStock(ctx context.Context, productID string, qty int) error
	DeductStock(ctx context.Context, productID string, qty int) error
	AcquireLock(ctx context.Context, orderID string, productID string, customerID string, qty int, ttlMinutes int) (lockID string, err error)
}

// OrderModel maps to the orders table
type OrderModel struct {
	OrderID     string
	CustomerID  string
	CartID      string
	Status      string
	TotalAmount float64
	CreatedAt   string
	UpdatedAt   string
}

// InventoryModel maps to the inventory table
type InventoryModel struct {
	InventoryID      string
	ProductID        string
	Quantity         int
	ReservedQuantity int
	LastUpdated      string
}
