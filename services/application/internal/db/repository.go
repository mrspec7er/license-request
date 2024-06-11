package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/mrspec7er/license-request-utility/dto"
	"gorm.io/gorm"
)

type Repository[T any, V any] interface {
	GetApp(V, string) error
	GetOne(T, string) error
	Create(T) error
	Update(T, string) error
	Delete(T, string) error
	GetAll(*[]T) error
}

type AppRepository struct {
	DB *Conn
}

func (r AppRepository) GetApp(app *dto.Form, number string) error {
	return r.DB.Preload("Sections.Fields.Responses", func(db *gorm.DB) *gorm.DB {
		return db.Where("responses.application_number = ?", number)
	}).First(&app).Error
}

func (r AppRepository) GetOne(app *dto.Application, number string) error {
	return r.DB.Where("number = ?", number).First(app).Error
}

func (r AppRepository) GetAll(apps *[]*dto.Application) error {
	return r.DB.Preload("User").Find(&apps).Error
}

func (r AppRepository) Create(app *dto.Application) error {
	return r.DB.Create(&app).Error
}

func (r AppRepository) Update(app *dto.Application, number string) error {
	return r.DB.Model(&dto.Application{}).Where("number = ?", number).Updates(&app).Error
}

func (r AppRepository) Delete(app *dto.Application, number string) error {
	return r.DB.Where("number = ?", app.Number).Delete(app).Error
}

type CacheRepository[T any] interface {
	Store(context.Context, string, T) error
	Retrieve(context.Context, string, T) error
}

type RedisRepository struct {
	Cache *CacheClient
}

func (r RedisRepository) Store(ctx context.Context, key string, value *dto.User) error {
	stringifiedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.Cache.Set(ctx, key, stringifiedValue, time.Hour*72).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r RedisRepository) Retrieve(ctx context.Context, key string, result *dto.User) error {
	value, err := r.Cache.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(value), &result)
	if err != nil {
		return err
	}

	return nil
}
