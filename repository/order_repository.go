package repository

import "context"

// OrderRepository defines database operations for Order
type OrderRepository interface {
	// Create inserts a new order record and returns it with generated ID
	Create(ctx context.Context, order *OrderModel) (*OrderModel, error)

	// FindByID retrieves an order by its ID
	FindByID(ctx context.Context, orderID string) (*OrderModel, error)

	// FindByCustomerID returns paginated orders for a customer
	// returns: orders, total count, error
	FindByCustomerID(ctx context.Context, customerID string, offset int, limit int) ([]*OrderModel, int, error)

	// UpdateStatus changes the status field of an order
	UpdateStatus(ctx context.Context, orderID string, status string) error
}

// InventoryRepository defines database operations for Inventory
type InventoryRepository interface {
	// FindByProductID retrieves inventory record for a product
	FindByProductID(ctx context.Context, productID string) (*InventoryModel, error)

	// ReserveStock atomically reserves stock — must use SELECT FOR UPDATE
	ReserveStock(ctx context.Context, productID string, qty int) error

	// ReleaseStock returns previously reserved stock back to available
	ReleaseStock(ctx context.Context, productID string, qty int) error

	// DeductStock permanently removes stock after order fulfillment
	DeductStock(ctx context.Context, productID string, qty int) error

	// AcquireLock saves a timed inventory lock record to database
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
