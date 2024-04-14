package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func StartConnection() *sqlx.DB {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT")))
	if err != nil {
		panic(err)
	}

	return db
}
