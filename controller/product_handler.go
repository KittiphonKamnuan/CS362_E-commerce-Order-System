package controller

import (
	"net/http"

	"github.com/cs362/ecommerce/dto/response"
	"github.com/cs362/ecommerce/service"
)

// ProductHandler handles HTTP requests for product endpoints
type ProductHandler struct {
	productService service.ProductService
}

// NewProductHandler creates a new ProductHandler
func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// ListProducts handles GET /api/v1/products
func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	page := queryInt(r, "page", 1)
	pageSize := queryInt(r, "pageSize", 10)

	products, err := h.productService.ListProducts(r.Context(), page, pageSize)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(products))
}

// GetProduct handles GET /api/v1/products/{productId}
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := extractPathParam(r.URL.Path, "products")

	product, err := h.productService.GetProductByID(r.Context(), productID)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(product))
}

// SearchProducts handles GET /api/v1/products/search?q=keyword
func (h *ProductHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("q")
	page := queryInt(r, "page", 1)
	pageSize := queryInt(r, "pageSize", 10)

	products, err := h.productService.SearchProducts(r.Context(), keyword, page, pageSize)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(products))
}
