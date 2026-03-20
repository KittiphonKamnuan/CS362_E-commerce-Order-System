package repository

import (
	"context"

	"github.com/cs362/ecommerce/entity"
)

// ProductRepository defines persistence operations for products
type ProductRepository interface {
	// FindByID retrieves a product by its ID
	FindByID(ctx context.Context, productID string) (*entity.Product, error)

	// FindAll returns paginated products
	FindAll(ctx context.Context, offset, limit int) ([]*entity.Product, int, error)

	// Search returns products matching a keyword in name or description
	Search(ctx context.Context, keyword string, offset, limit int) ([]*entity.Product, int, error)
}
