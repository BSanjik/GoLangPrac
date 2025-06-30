package storage

import (
	"go_shop/model"
)

func (storage *ProductStorage) ReturnAllProducts() ([]model.Product, error) {
	//SQL request POSTGre

	//give data from db
	return storage.data, nil
}

func (storage *ProductStorage) CreateProduct(newProduct model.Product) error {
	storage.lastID++
	newProduct.ID = storage.lastID

	storage.data = append(storage.data, newProduct)
	return nil
}
