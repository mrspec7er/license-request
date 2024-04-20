package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/mrspec7er/license-request/service/user/internal/db"
	"gorm.io/gorm"
)

type key string

const (
	UserContextKey key = "user"
)

func AuthInit() *sessions.CookieStore {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	maxAge := 86400 * 1 // 1 days
	isProd := false     // Set to true when serving over https

	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_AUTH_KEY"), os.Getenv("GOOGLE_AUTH_SECRET"), os.Getenv("API_URL")+"/auth/callback?provider=google"),
	)

	return store
}

type AuthService struct {
	DB    *gorm.DB
	Store *sessions.CookieStore
}

func (s AuthService) GetUserEmail(r *http.Request, userEmail *string) error {
	session, err := s.Store.Get(r, "auth")
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

func (s AuthService) FindUser(user *db.User) (int, error) {
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

	session, err := s.Store.Get(r, "auth")
	if err != nil {
		return nil, err
	}
	session.Values["email"] = user.Email
	session.Save(r, w)

	return &user, nil
}

func (s AuthService) RemoveUserSessions(w http.ResponseWriter, r *http.Request) {
	session, _ := s.Store.Get(r, "auth")
	session.Values["email"] = nil
	session.Save(r, w)
	gothic.Logout(w, r)
}
