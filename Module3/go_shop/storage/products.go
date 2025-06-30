package storage

import "go_shop/model"

func (storage *ProductStorage) ReturnAllProducts() ([]model.Product, error) {
	//SQL request POSTGre

	//give data from db
	return storage.data, nil
}
