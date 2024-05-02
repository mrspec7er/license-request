package src

import (
	"github.com/mrspec7er/license-request-utility/dto"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func (s Service) GetOne(form *dto.Form, id uint) (int, error) {
	err := s.DB.Preload("Sections").Preload("Sections.Fields").First(&form, id).Error

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s Service) Create(form *dto.Form) (int, error) {
	err := s.DB.Save(&form).Error

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s Service) Delete(form *dto.Form) (int, error) {
	err := s.DB.Delete(&dto.Form{}, form.ID).Error

	if err != nil {
		return 500, err
	}

	return 200, nil
}
