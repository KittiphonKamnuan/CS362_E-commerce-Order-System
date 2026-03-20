package entity

import (
	"errors"
	"time"
)

// Inventory tracks stock for a product
type Inventory struct {
	InventoryID      string
	Product          *Product
	Quantity         int
	ReservedQuantity int
	LastUpdated      time.Time
}

// GetAvailableStock returns stock available for new orders
func (inv *Inventory) GetAvailableStock() int {
	return inv.Quantity - inv.ReservedQuantity
}

// ReserveStock temporarily holds stock for a pending order
func (inv *Inventory) ReserveStock(qty int) error {
	if qty <= 0 {
		return errors.New("quantity must be positive")
	}
	if inv.GetAvailableStock() < qty {
		return errors.New("insufficient stock")
	}
	inv.ReservedQuantity += qty
	inv.LastUpdated = time.Now()
	return nil
}

// ReleaseStock frees previously reserved stock
func (inv *Inventory) ReleaseStock(qty int) error {
	if qty <= 0 {
		return errors.New("quantity must be positive")
	}
	if inv.ReservedQuantity < qty {
		return errors.New("cannot release more than reserved quantity")
	}
	inv.ReservedQuantity -= qty
	inv.LastUpdated = time.Now()
	return nil
}

// DeductStock permanently removes stock after order fulfillment
func (inv *Inventory) DeductStock(qty int) error {
	if qty <= 0 {
		return errors.New("quantity must be positive")
	}
	if inv.Quantity < qty {
		return errors.New("insufficient stock for deduction")
	}
	inv.Quantity -= qty
	if inv.ReservedQuantity >= qty {
		inv.ReservedQuantity -= qty
	}
	inv.LastUpdated = time.Now()
	return nil
}

// InventoryLock represents a time-limited hold on inventory for an order
type InventoryLock struct {
	LockID     string
	OrderID    string
	ProductID  string
	CustomerID string
	Quantity   int
	ExpiresAt  time.Time
}

// IsExpired returns true if the lock has passed its expiry time
func (l *InventoryLock) IsExpired() bool {
	return time.Now().After(l.ExpiresAt)
}
