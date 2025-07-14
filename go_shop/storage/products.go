package storage

import (
	"context"
	"fmt"
	"go_shop/model"
)

func (storage *ProductStorage) ReturnAllProducts() ([]model.Product, error) {
	//SQL request POSTGre
	query := `SELECT * FROM products`

	result, err := storage.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var products []model.Product
	for result.Next() {
		var prod model.Product
		err = result.Scan(&prod.ID, &prod.Name, &prod.Stock, &prod.Price)
		if err != nil {
			return nil, fmt.Errorf("error parsing %v", err)
		}
		products = append(products, prod)
	}

	//give data from db
	return products, nil
}

func (storage *ProductStorage) CreateProduct(newProduct model.Product) error {
	query := `
		insert into PRODUCTS (name, stock, price)
		VALUES ($1, $2, $3)
	`
	_, err := storage.DB.Exec(context.Background(), query, newProduct.Name, newProduct.Stock, newProduct.Price)
	if err != nil {
		return err
	}
	return nil
}

// func (storage *ProductStorage) BuyProduct(name string, quantity int) (float64, error) {
// 	for i := range storage.data {
// 		if storage.data[i].Name == name {
// 			if storage.data[i].Stock < quantity {
// 				return 0, fmt.Errorf("not enough stock for %s: have %d, need %d", name, storage.data[i].Stock, quantity)
// 			}
// 			storage.data[i].Stock -= quantity
// 			return storage.data[i].Price, nil
// 		}
// 	}
// 	return 0, fmt.Errorf("product %s not found", name)
// }
