package storage

import "go_shop/model"

type ProductStorage struct {
	data   []model.Product
	lastID int
}

var products []model.Product = []model.Product{
	{
		ID:    1,
		Name:  "Bread",
		Stock: 2,
		Price: 50.0,
	},
	{
		ID:    2,
		Name:  "Soap",
		Stock: 3,
		Price: 150.0,
	},
}

func NewProductStorage() *ProductStorage {
	return &ProductStorage{
		data:   products,
		lastID: len(products),
	}
}
