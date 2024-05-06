package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mrspec7er/license-request/services/logger/internal"
	"github.com/mrspec7er/license-request/services/logger/internal/db"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	DB := db.StartConnection()
	Memcache := db.MemcacheConnection()

	config := &internal.Server{
		DB:       DB,
		Memcache: Memcache,
	}

	server := internal.NewServer(*config)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
