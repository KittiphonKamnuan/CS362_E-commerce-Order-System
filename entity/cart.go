package entity

import (
	"errors"
	"time"
)

// CartItem represents a single product entry in a cart
type CartItem struct {
	CartItemID string
	Product    *Product
	Quantity   int
	UnitPrice  float64
}

// GetSubtotal returns the total price for this cart item
func (ci *CartItem) GetSubtotal() float64 {
	return ci.UnitPrice * float64(ci.Quantity)
}

// Cart holds items selected by a customer before checkout
type Cart struct {
	CartID     string
	Customer   *Customer
	Items      []*CartItem
	CreatedAt  time.Time
	UpdatedAt  time.Time
	TotalPrice float64
}

// CalculateTotal sums all cart item subtotals
func (c *Cart) CalculateTotal() float64 {
	total := 0.0
	for _, item := range c.Items {
		total += item.GetSubtotal()
	}
	c.TotalPrice = total
	return total
}

// IsValid returns an error if the cart cannot be checked out
func (c *Cart) IsValid() error {
	if len(c.Items) == 0 {
		return errors.New("cart is empty")
	}
	return nil
}

// Clear resets the cart to an empty state
func (c *Cart) Clear() {
	c.Items = []*CartItem{}
	c.TotalPrice = 0
	c.UpdatedAt = time.Now()
}
