package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ProductID    int32  `json:"product_id"`
	SKU          string `json:"sku"`
	Manufacturer string `json:"manufacturer"`
	CategoryID   int32  `json:"category_id"`
	Weight       int32  `json:"weight"`
	SomeOtherID  int32  `json:"some_other_id"`
}

var products = make(map[int32]Product)

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	// Expect paths like:
	//   /products/{id}           (GET)
	//   /products/{id}/details   (POST)
	path := strings.TrimPrefix(r.URL.Path, "/products/")
	if path == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	// detect /details suffix
	isDetails := strings.HasSuffix(path, "/details")
	idStr := path
	if isDetails {
		idStr = strings.TrimSuffix(path, "/details")
	}
	idStr = strings.Trim(idStr, "/")

	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}
	productID := int32(id64)

	switch r.Method {
	case http.MethodGet:
		if isDetails {
			http.Error(w, "GET not allowed on /details", http.StatusMethodNotAllowed)
			return
		}
		getProduct(w, r, productID)
	case http.MethodPost:
		if !isDetails {
			http.Error(w, "POST must target /details", http.StatusBadRequest)
			return
		}
		addProductDetails(w, r, productID)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getProduct(w http.ResponseWriter, r *http.Request, productID int32) {
	log.Printf("GET /products/%d", productID)
	product, ok := products[productID]
	if !ok {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func addProductDetails(w http.ResponseWriter, r *http.Request, productID int32) {
	log.Printf("POST /products/%d/details", productID)
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	products[productID] = newProduct
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/products/", productsHandler)

	port := ":8081" // binds on 0.0.0.0 by default
	log.Printf("Server starting on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
