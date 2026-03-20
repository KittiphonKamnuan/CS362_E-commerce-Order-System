package response

import "time"

// ProductResponse is the DTO returned for product queries
type ProductResponse struct {
	ProductID   string    `json:"productId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	ImageURL    string    `json:"imageUrl"`
	Status      string    `json:"status"`
	CategoryID  string    `json:"categoryId"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"createdAt"`
}

// ProductListResponse wraps a paginated list of products
type ProductListResponse struct {
	Products []*ProductResponse `json:"products"`
	Total    int                `json:"total"`
	Page     int                `json:"page"`
	PageSize int                `json:"pageSize"`
}
