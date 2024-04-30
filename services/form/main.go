package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mrspec7er/license-request/services/form/internal"
	"github.com/mrspec7er/license-request/services/form/internal/db"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	DB := db.StartConnection()

	config := &internal.Server{
		DB: DB,
	}

	dbConn, err := DB.DB()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
	defer dbConn.Close()

	server := internal.NewServer(*config)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
