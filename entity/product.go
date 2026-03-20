package entity

import "time"

// Category groups related products
type Category struct {
	CategoryID  string
	Name        string
	Description string
}

// Product represents a sellable item in the store
type Product struct {
	ProductID   string
	Name        string
	Description string
	Price       float64
	ImageURL    string
	Status      ProductStatus
	CreatedAt   time.Time
	Category    *Category
}

// CheckAvailability returns true if the product is active
func (p *Product) CheckAvailability() bool {
	return p.Status == ProductStatusActive
}
