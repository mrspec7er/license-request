package application

import (
	"github.com/mrspec7er/license-request/server/internal/db"
	"gorm.io/gorm"
)

type ApplicationService struct {
	DB *gorm.DB
}

func (s ApplicationService) GetOne(app *db.Form, number string) (int, error) {
	err := s.DB.Preload("Sections.Fields.Responses", func(db *gorm.DB) *gorm.DB {
		return db.Where("responses.application_number = ?", number)
	}).First(&app).Error

	if err != nil {
		return 500, err
	}

	return 200, nil
}
