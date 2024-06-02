package db

import (
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
