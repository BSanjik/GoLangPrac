package storage

import (
	"fmt"
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

func (storage *ProductStorage) BuyProduct(name string, quantity int) (float64, error) {
	for i := range storage.data {
		if storage.data[i].Name == name {
			if storage.data[i].Stock < quantity {
				return 0, fmt.Errorf("not enough stock for %s: have %d, need %d", name, storage.data[i].Stock, quantity)
			}
			storage.data[i].Stock -= quantity
			return storage.data[i].Price, nil
		}
	}
	return 0, fmt.Errorf("product %s not found", name)
}
