package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	//empty context
	ctx := context.Background()

	// login to db /creds
	dsn := "postgres://postgres:postgres@localhost:5432/test_db" //postgres://go_user:go_pass@localhost:5432/test_db

	//connect to db
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx) //close of context

	log.Println("Connection successfully")

	//sql to db
}
