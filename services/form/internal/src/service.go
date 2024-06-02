package src

import (
	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/mrspec7er/license-request/services/form/internal/db"
)

type Service struct {
	Store db.Repository[*dto.Form]
}

func (s Service) GetOne(form *dto.Form, id uint) (int, error) {
	err := s.Store.GetOne(form, id)

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s Service) GetAll(form *[]*dto.Form) (int, error) {
	err := s.Store.GetAll(form)

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s Service) Create(form *dto.Form) (int, error) {
	err := s.Store.Create(form)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s Service) Delete(form *dto.Form) (int, error) {
	err := s.Store.Delete(&dto.Form{}, form.ID)

	if err != nil {
		return 500, err
	}

	return 200, nil
}
