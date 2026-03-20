package service

import (
	"context"

	"github.com/cs362/ecommerce/dto/response"
)

// ProductService defines business logic for product catalog
type ProductService interface {
	// GetProductByID retrieves a single product by ID
	GetProductByID(ctx context.Context, productID string) (*response.ProductResponse, error)

	// ListProducts returns paginated products
	ListProducts(ctx context.Context, page, pageSize int) (*response.ProductListResponse, error)

	// SearchProducts returns products matching a keyword
	SearchProducts(ctx context.Context, keyword string, page, pageSize int) (*response.ProductListResponse, error)
}
