package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mrspec7er/license-request/server/internal"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT")))
	if err != nil {
		panic(err)
	}

	config := &internal.Server{
		DB: db,
	}

	server := internal.NewServer(*config)
	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
