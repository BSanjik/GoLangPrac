package handler

import (
	"encoding/json"
	"fmt"
	"go_shop/model"
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
	case r.Method == http.MethodPost && r.URL.Path == "/products":
		h.CreateProduct(w, r)
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
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error")
		return
	}

	//take response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// Create Product
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct model.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		log.Println("Error parsing", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Check products")
		return
	}
	defer r.Body.Close()

	err = h.Service.CreateProduct(newProduct)
	if err != nil {
		log.Println("Error: ", err)
		//return response to user
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error")
		return
	}

	//take response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Success")
}
