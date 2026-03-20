// Package main is the entry point for the E-commerce Order System API.
//
//	@title			CS362 E-commerce Order System API
//	@version		1.0
//	@description	Flash Sale E-commerce API with Rate Limiting — CS362 Final Project
//	@termsOfService	http://swagger.io/terms/
//
//	@contact.name	CS362 Team
//
//	@host		localhost:8080
//	@BasePath	/api/v1
//
//	@securityDefinitions.apikey	CustomerID
//	@in							header
//	@name						X-Customer-ID
//	@description				Customer identifier passed as request header
package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/cs362/ecommerce/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	mux := http.NewServeMux()

	// Swagger UI
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"status":"ok"}`)
	})

	// Order endpoints
	// POST   /api/v1/orders
	// GET    /api/v1/orders/{orderId}
	// GET    /api/v1/orders
	// PATCH  /api/v1/orders/{orderId}/cancel

	// Product endpoints
	// GET    /api/v1/products
	// GET    /api/v1/products/{productId}
	// GET    /api/v1/products/search

	// Cart endpoints
	// GET    /api/v1/cart
	// POST   /api/v1/cart/items
	// PATCH  /api/v1/cart/items/{productId}
	// DELETE /api/v1/cart/items/{productId}

	// Payment endpoints
	// POST   /api/v1/payments
	// POST   /api/v1/payments/{paymentId}/refund

	addr := ":8080"
	log.Printf("Server starting on %s", addr)
	log.Printf("Swagger UI: http://localhost%s/swagger/index.html", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
