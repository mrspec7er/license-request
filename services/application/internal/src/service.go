package src

import (
	"github.com/mrspec7er/license-request/services/utility/dto"
	"gorm.io/gorm"
)

type ApplicationService struct {
	DB   *gorm.DB
	Util *ApplicationUtil
}

func (s ApplicationService) GetOne(app *dto.Form, number string) (int, error) {
	err := s.DB.Preload("Sections.Fields.Responses", func(db *gorm.DB) *gorm.DB {
		return db.Where("responses.application_number = ?", number)
	}).First(&app).Error

	if err != nil {
		return 500, err
	}

	return 200, nil
}
