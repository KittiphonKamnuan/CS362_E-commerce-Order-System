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

// ListProducts godoc
//
//	@Summary		List all products
//	@Description	Returns a paginated list of available products
//	@Tags			Products
//	@Produce		json
//	@Param			page		query		int	false	"Page number"	default(1)
//	@Param			pageSize	query		int	false	"Page size"		default(10)
//	@Success		200			{object}	response.SuccessResponse{data=response.ProductListResponse}
//	@Failure		500			{object}	response.ErrorResponse
//	@Router			/products [get]
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

// GetProduct godoc
//
//	@Summary		Get product by ID
//	@Description	Retrieves a single product by its ID
//	@Tags			Products
//	@Produce		json
//	@Param			productId	path		string	true	"Product ID"
//	@Success		200			{object}	response.SuccessResponse{data=response.ProductResponse}
//	@Failure		404			{object}	response.ErrorResponse
//	@Failure		500			{object}	response.ErrorResponse
//	@Router			/products/{productId} [get]
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := extractPathParam(r.URL.Path, "products")

	product, err := h.productService.GetProductByID(r.Context(), productID)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response.NewSuccess(product))
}

// SearchProducts godoc
//
//	@Summary		Search products
//	@Description	Returns products matching a search keyword in name or description
//	@Tags			Products
//	@Produce		json
//	@Param			q			query		string	true	"Search keyword"
//	@Param			page		query		int		false	"Page number"	default(1)
//	@Param			pageSize	query		int		false	"Page size"		default(10)
//	@Success		200			{object}	response.SuccessResponse{data=response.ProductListResponse}
//	@Failure		500			{object}	response.ErrorResponse
//	@Router			/products/search [get]
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
