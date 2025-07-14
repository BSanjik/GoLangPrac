package storage

import (
	"github.com/jackc/pgx/v5"
)

type ProductStorage struct {
	DB *pgx.Conn
}

// var products []model.Product = []model.Product{
// 	{
// 		ID:    1,
// 		Name:  "Bread",
// 		Stock: 2,
// 		Price: 50.0,
// 	},
// 	{
// 		ID:    2,
// 		Name:  "Soap",
// 		Stock: 3,
// 		Price: 150.0,
// 	},
// }

func NewProductStorage(db *pgx.Conn) *ProductStorage {
	return &ProductStorage{
		DB: db,
	}
}
