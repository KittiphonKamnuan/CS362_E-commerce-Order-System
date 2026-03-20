package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Routes are registered here once service/repository implementations are provided.
	// Example route stubs:

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

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"status":"ok"}`)
	})

	addr := ":8080"
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
