package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mrspec7er/license-request/server/internal"
	"github.com/mrspec7er/license-request/server/internal/db"
	"github.com/mrspec7er/license-request/server/internal/util"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	util.AuthInit()
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
