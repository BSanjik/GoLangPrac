package service

import (
	"go_shop/model"
	"go_shop/storage"
	"log"
)

type ProductService struct {
	Storage *storage.ProductStorage
}

func NewProductService(storage *storage.ProductStorage) *ProductService {
	return &ProductService{Storage: storage}
}

func (s *ProductService) ReturnAllProducts() ([]model.Product, error) {
	//give handler product list
	result, err := s.Storage.ReturnAllProducts()
	if err != nil {
		log.Println("Error storage: ", err)
		return nil, err
	}

	return result, nil
}

func (s *ProductService) CreateProduct(newProduct model.Product) error {
	return s.Storage.CreateProduct(newProduct)
}
