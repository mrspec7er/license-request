package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	CreateDB()
}

func CreateDB() *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable TimeZone=Asia/Singapore", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	cmd := fmt.Sprintf("SELECT 'CREATE DATABASE %s' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '%s')", os.Getenv("DB_NAME"), os.Getenv("DB_NAME"))
	db.Exec(cmd)

	fmt.Println("Database created!")

	return db
}
