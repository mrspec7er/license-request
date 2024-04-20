package form

import (
	"github.com/mrspec7er/license-request/user/internal/db"
	"gorm.io/gorm"
)

type FormService struct {
	DB *gorm.DB
}

func (s FormService) GetOne(form *db.Form) (int, error) {
	err := s.DB.Preload("Sections").Preload("Sections.Fields").First(&form).Error

	if err != nil {
		return 500, err
	}

	return 200, nil
}
