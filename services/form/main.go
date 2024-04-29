package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mrspec7er/license-request/services/form/internal"
	"github.com/mrspec7er/license-request/services/form/internal/db"
	"github.com/mrspec7er/license-request/services/form/internal/hub"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	DB := db.StartConnection()
	Hub := hub.StartConnection()

	config := &internal.Server{
		DB:  DB,
		Hub: Hub,
	}

	dbConn, err := DB.DB()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	defer Hub.Close()
	defer dbConn.Close()

	server := internal.NewServer(*config)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
