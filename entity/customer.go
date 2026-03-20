package entity

import "time"

// Address represents a physical address
type Address struct {
	AddressID  string
	Street     string
	City       string
	Province   string
	State      string
	PostalCode string
	Country    string
}

// GetFullAddress returns a formatted full address string
func (a *Address) GetFullAddress() string {
	return a.Street + ", " + a.City + ", " + a.Province + ", " + a.PostalCode + ", " + a.Country
}

// Customer represents a registered user of the e-commerce system
type Customer struct {
	CustomerID string
	Name       string
	Email      string
	Phone      string
	Addresses  []*Address
	Password   string
	CreatedAt  time.Time
}
