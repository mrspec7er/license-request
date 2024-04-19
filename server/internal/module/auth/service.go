package auth

import (
	"fmt"
	"net/http"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/mrspec7er/license-request/server/internal/db"
	"github.com/mrspec7er/license-request/server/internal/util"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func (s AuthService) GetUserEmail(r *http.Request, userEmail *string) error {
	session, err := util.Store.Get(r, "auth")
	if err != nil {
		return err
	}
	email, ok := session.Values["email"].(string)

	if !ok || email == "" {
		return fmt.Errorf("unauthorize user")
	}

	*userEmail = email

	return nil
}

func (s AuthService) FindUser(user *db.Account) (int, error) {
	err := s.DB.Where("email = ?", user.Email).First(&user).Error

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s AuthService) SaveUserSessions(w http.ResponseWriter, r *http.Request) (*goth.User, error) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return nil, err
	}

	session, err := util.Store.Get(r, "auth")
	if err != nil {
		return nil, err
	}
	session.Values["email"] = user.Email
	session.Save(r, w)

	return &user, nil
}

func (s AuthService) RemoveUserSessions(w http.ResponseWriter, r *http.Request) {
	session, _ := util.Store.Get(r, "auth")
	session.Values["email"] = nil
	session.Save(r, w)
	gothic.Logout(w, r)
}
