package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/mrspec7er/license-request-utility/dto"
)

type Repository[T any] interface {
	GetOne(T, uint) error
	Create(T) error
	Delete(T, uint) error
	GetAll(*[]T) error
}

type FormRepository struct {
	DB *Conn
}

func (r FormRepository) GetOne(form *dto.Form, id uint) error {
	return r.DB.Preload("Sections.Fields").First(&form, id).Error
}

func (r FormRepository) GetAll(forms *[]*dto.Form) error {
	return r.DB.Preload("Sections.Fields").Find(&forms).Error
}

func (r FormRepository) Create(form *dto.Form) error {
	return r.DB.Create(&form).Error
}

func (r FormRepository) Delete(form *dto.Form, id uint) error {
	return r.DB.Delete(&form, id).Error
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
