package entity

import (
	"errors"
	"time"
)

// OrderItem represents a product line within an order
type OrderItem struct {
	Product   *Product
	Quantity  int
	UnitPrice float64
}

// GetSubtotal returns the total price for this order item
func (oi *OrderItem) GetSubtotal() float64 {
	return oi.UnitPrice * float64(oi.Quantity)
}

// Order represents a customer's purchase request
type Order struct {
	OrderID           string
	Customer          *Customer
	Items             []*OrderItem
	Status            OrderStatus
	TotalAmount       float64
	ShippingAddressID string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// CalculateTotal sums all order item subtotals
func (o *Order) CalculateTotal() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += item.GetSubtotal()
	}
	o.TotalAmount = total
	return total
}

// validTransitions defines allowed order state machine moves
var validTransitions = map[OrderStatus][]OrderStatus{
	OrderStatusPending:    {OrderStatusConfirmed, OrderStatusCancelled},
	OrderStatusConfirmed:  {OrderStatusProcessing, OrderStatusCancelled},
	OrderStatusProcessing: {OrderStatusShipped},
	OrderStatusShipped:    {OrderStatusDelivered},
	OrderStatusDelivered:  {OrderStatusRefunded},
	OrderStatusCancelled:  {},
	OrderStatusRefunded:   {},
}

// UpdateStatus transitions the order to a new status, enforcing state machine rules
func (o *Order) UpdateStatus(status OrderStatus) error {
	allowed, ok := validTransitions[o.Status]
	if !ok {
		return errors.New("unknown current status")
	}
	for _, s := range allowed {
		if s == status {
			o.Status = status
			o.UpdatedAt = time.Now()
			return nil
		}
	}
	return errors.New("invalid status transition from " + string(o.Status) + " to " + string(status))
}

// CanBeCancelled returns true if the order is in a cancellable state
func (o *Order) CanBeCancelled() bool {
	return o.Status == OrderStatusPending || o.Status == OrderStatusConfirmed
}

// IsValid returns an error if the order is missing required fields
func (o *Order) IsValid() error {
	if o.Customer == nil || o.Customer.CustomerID == "" {
		return errors.New("order must have a customer")
	}
	if len(o.Items) == 0 {
		return errors.New("order must have at least one item")
	}
	if o.TotalAmount <= 0 {
		return errors.New("order total amount must be positive")
	}
	return nil
}
