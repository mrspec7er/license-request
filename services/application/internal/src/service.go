package src

import (
	"errors"

	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/mrspec7er/license-request/services/application/internal/db"
)

type Service struct {
	Store db.Repository[*dto.Application, *dto.Form]
	Util  *Util
}

func (s Service) GetOne(app *dto.Form, number string) (int, error) {
	err := s.Store.GetApp(app, number)
	if err != nil {
		return 400, err
	}

	return 200, nil
}

func (s Service) GetAll(apps *[]*dto.Application, number string) (int, error) {
	err := s.Store.GetAll(apps)

	if err != nil {
		return 400, err
	}

	return 200, nil
}

func (s Service) Create(app *dto.Application) (int, error) {
	app.Status = string(dto.RequestNew)
	err := s.Store.Create(app)

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s Service) ChangeStatus(number string, status string, note string, uid string) (int, error) {
	if status != string(dto.RequestApproved) && status != string(dto.RequestRejected) {
		return 400, errors.New("invalid status type")
	}

	app := &dto.Application{Status: status, ApprovedBy: uid, Note: note}

	err := s.Store.Update(app, number)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s Service) Delete(app *dto.Application) (int, error) {
	err := s.Store.Delete(app, app.Number)

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s Service) ApplicationAccessGuard(number string, user dto.User) (int, error) {
	app := &dto.Application{}
	err := s.Store.GetOne(app, number)

	if err != nil {
		return 400, err
	}

	if app.UserID != user.ID && user.Role != "ADMIN" {
		return 403, errors.New("user access denied")
	}

	return 200, nil
}
