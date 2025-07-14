package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func main() {
	//empty context
	ctx := context.Background()

	// login to db /creds
	dsn := "postgres://go_user:go_pass@localhost:5434/test_db?sslmode=disable" //postgres://go_user:go_pass@localhost:5432/test_db

	//connect to db
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx) //close of context

	log.Println("Connection successfully")

	//sql to db
	QueryCreateCats := `
	CREATE TABLE IF NOT EXISTS cats(
            id UUID PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            color VARCHAR(255) NOT NULL,
            weight INT NULL,
            breed VARCHAR(255) NOT NULL,
            sex CHAR(1) NOT NULL
        )
	`
	_, err = conn.Exec(ctx, QueryCreateCats)
	if err != nil {
		log.Fatalf("ERROR")
	}
	log.Println("Cats created")

	QueryInsertCats := `
		INSERT INTO cats(id, name, color, weight, breed, sex)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err = conn.Exec(ctx, QueryInsertCats, uuid.New(), "Barsik", "Black", 4, "NoBreed", "F")
	if err != nil {
		log.Printf("Error inserting cat : %v", err)
	} else {
		log.Println("Cat created")
	}
}
