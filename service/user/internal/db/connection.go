package db

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartConnection() *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=Asia/Singapore", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&User{},
	)

	return db
}

func MemcacheConnection() *redis.Client {
	memcache := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_USERNAME"),
		Username: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return memcache
}
