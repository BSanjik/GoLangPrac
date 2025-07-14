package main

import (
	"context"
	"go_shop/handler"
	"go_shop/service"
	"go_shop/storage"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	dsn := "postgres://go_user:go_pass@localhost:5434/test_db"

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("Error connecting DB: %v", err)
	}
	defer conn.Close(ctx)

	//создвет объекты хранилища с товарами
	productStorage := storage.NewProductStorage(conn)
	productService := service.NewProductService(productStorage)
	h := handler.NewProductHandler(productService)

	//все запросы от клиента фронт
	mux := http.NewServeMux()
	mux.Handle("/products", h)
	mux.Handle("/products/buy", h)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server is started on port :8080")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error:%v", err)
	}
}

//curl request
// curl -X POST http://localhost:7777/products \
// -H "Content-Type: application/json" \
// -d '{"name": "БМВ М5", "stock": 1000, "price": 85000000.0}'
