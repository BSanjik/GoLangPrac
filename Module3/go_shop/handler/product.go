package handler

import (
	"encoding/json"
	"fmt"
	"go_shop/service"
	"log"
	"net/http"
)

type ProductHandler struct {
	Service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{Service: s}
}

// Main router
func (h *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/products":
		h.GetAllProducts(w, r)
	default:
		http.NotFound(w, r)
	}
}

// Get All Products
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	//call service

	result, err := h.Service.ReturnAllProducts()
	if err != nil {
		//log error
		log.Printf("Error: %v", err)
		//return response to user
		fmt.Fprint(w, "Error")
		return
	}

	//take response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
