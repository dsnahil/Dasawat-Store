package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Product defines the structure for a product based on the api.yaml.
type Product struct {
	ProductID    int32  `json:"product_id"`
	SKU          string `json:"sku"`
	Manufacturer string `json:"manufacturer"`
	CategoryID   int32  `json:"category_id"`
	Weight       int32  `json:"weight"`
	SomeOtherID  int32  `json:"some_other_id"`
}

// A simple in-memory database using a map.
// The key is the productID (int32).
var products = make(map[int32]Product)

// productsHandler is the primary router for all /products/ requests.
func productsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the productID from the URL path.
	// e.g., /products/123 -> "123"
	// e.g., /products/123/details -> "123"
	idStr := strings.TrimPrefix(r.URL.Path, "/products/")
	idStr = strings.TrimSuffix(idStr, "/details")

	// If idStr is empty, it means the request is to /products/ which is not a valid endpoint.
	if idStr == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	// Convert the productID from string to int32.
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}
	productID := int32(id)

	// Route the request based on the HTTP method.
	switch r.Method {
	case http.MethodGet:
		getProduct(w, r, productID)
	case http.MethodPost:
		addProductDetails(w, r, productID)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getProduct handles GET /products/{productId}
func getProduct(w http.ResponseWriter, r *http.Request, productID int32) {
	log.Printf("Received GET request for product ID: %d", productID)

	// Look for the productID in the 'products' map.
	product, ok := products[productID]
	if !ok {
		// If 'ok' is false, the product was not found. Return a 404 error.
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// If the product is found, encode it to JSON and send it back.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// addProductDetails handles POST /products/{productId}/details
func addProductDetails(w http.ResponseWriter, r *http.Request, productID int32) {
	log.Printf("Received POST request for product ID: %d", productID)

	var newProduct Product

	// Decode the JSON request body into the 'newProduct' variable.
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		// If there's an error, the request body was invalid.
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Store the new product details in our map using the ID from the URL.
	products[productID] = newProduct
	log.Printf("Added/Updated details for product: %+v", newProduct)

	// Send a 204 No Content response to indicate success without a body.
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	// This router will send all requests starting with "/products/" to our handler.
	http.HandleFunc("/products/", productsHandler)

	port := ":8080"
	log.Printf("Server starting on port %s", port)

	// Start the HTTP server.
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
