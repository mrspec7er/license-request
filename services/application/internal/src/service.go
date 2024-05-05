package src

import (
	"errors"

	"github.com/mrspec7er/license-request-utility/dto"
	"gorm.io/gorm"
)

type Service struct {
	DB   *gorm.DB
	Util *ApplicationUtil
}

func (s Service) GetOne(app *dto.Form, number string) (int, error) {
	err := s.DB.Preload("Sections.Fields.Responses", func(db *gorm.DB) *gorm.DB {
		return db.Where("responses.application_number = ?", number)
	}).First(&app).Error

	if err != nil {
		return 400, err
	}

	return 200, nil
}

func (s Service) Create(app *dto.Application) (int, error) {
	app.Status = string(dto.RequestNew)
	err := s.DB.Create(&app).Error

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s Service) ChangeStatus(number string, status string, note string, uid string) (int, error) {
	if status != string(dto.RequestApproved) && status != string(dto.RequestRejected) {
		return 400, errors.New("invalid status type")
	}

	err := s.DB.Model(&dto.Application{}).Where("number = ?", number).Updates(&dto.Application{Status: status, ApprovedBy: uid, Note: note}).Error
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s Service) Delete(app *dto.Application) (int, error) {
	err := s.DB.Where("number = ?", app.Number).Delete(app).Error

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s Service) ApplicationAccessGuard(number string, user dto.User) (int, error) {
	app := &dto.Application{}
	err := s.DB.Where("number = ?", number).First(app).Error

	if err != nil {
		return 400, err
	}

	if app.UserID != user.ID && user.Role != "ADMIN" {
		return 403, errors.New("user access denied")
	}

	return 200, nil
}
