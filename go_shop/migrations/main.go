package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

func runMigrations(dsn string) {
	migr, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatalf("error migrating: %v", err)
	}

	_, _, err = migr.Version()
	if err == migrate.ErrNilVersion {
		log.Println("No migrations")
	} else if err != nil {
		log.Printf("Some error: %v", err)
	}

	err = migr.Up()
	if err != nil {
		if err.Error() != "no change" {
			log.Fatalf("все плохо: %v", err)
		}
	}
	fmt.Println("It's ok!!!")
}

func main() {
	dsn := "postgres://go_user:go_pass@localhost:5434/test_db?sslmode=disable"
	// 1. Проведение миграций
	runMigrations(dsn)

	// 2. Получение данных (чтобы проверить)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("error connecting to db: %v", conn)
	}
	defer conn.Close(context.Background())

	log.Println("Connected to DB successfully")

	query := `
        SELECT * FROM cats WHERE name = $1
    `

	var cat Cat
	err = conn.QueryRow(context.Background(), query, "Barsik").Scan(&cat.Id, &cat.Name)
	if err != nil {
		log.Fatalf("can not get cat from table: %v", err)
	}
	log.Printf("My cat: %v", cat)
}

type Cat struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
