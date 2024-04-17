package form

import (
	"github.com/mrspec7er/license-request/server/internal/dto"
	"gorm.io/gorm"
)

type FormService struct {
	DB *gorm.DB
}

func (s FormService) GetAll(form *[]*dto.Form) (int, error) {
	var err error

	if err != nil {
		return 500, err
	}

	return 200, nil
}
