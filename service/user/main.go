package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mrspec7er/license-request/user/internal"
	"github.com/mrspec7er/license-request/user/internal/db"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	db := db.StartConnection()

	config := &internal.Server{
		DB: db,
	}

	server := internal.NewServer(*config)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
