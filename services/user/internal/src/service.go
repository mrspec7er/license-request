package src

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/mrspec7er/license-request-utility/dto"
	"gorm.io/gorm"
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

type Service struct {
	DB    *gorm.DB
	Store *sessions.CookieStore
	Util  *Utility
}

func (s Service) GetUserEmail(r *http.Request, userEmail *string) error {
	session, err := s.Store.Get(r, string(dto.AuthCookieName))
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

func (s Service) FindUser(user *dto.User) (int, error) {
	err := s.DB.Where("email = ?", user.Email).First(&user).Error

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s Service) SaveUserSessions(w http.ResponseWriter, r *http.Request) (*goth.User, error) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return nil, err
	}

	session, err := s.Store.Get(r, string(dto.AuthCookieName))
	if err != nil {
		return nil, err
	}
	session.Values["email"] = user.Email
	session.Save(r, w)

	return &user, nil
}

func (s Service) StoreUserSessions(w http.ResponseWriter, r *http.Request, user *goth.User) error {
	data, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return err
	}

	key := data.UserID + "-" + time.Now().Local().String()

	err = s.Util.Store(context.Background(), key, data.RawData)
	if err != nil {
		return err
	}

	cookie := http.Cookie{
		Name:     string(dto.AuthCookieName),
		Value:    key,
		Path:     "/",
		MaxAge:   7200,
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	*user = data

	return nil
}

func (s Service) RetrieveUserSessions(w http.ResponseWriter, r *http.Request, key string, user *dto.User) error {

	err := s.Util.Retrieve(context.Background(), key, &user)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) RemoveUserSessions(w http.ResponseWriter, r *http.Request) {
	session, _ := s.Store.Get(r, string(dto.AuthCookieName))
	session.Values["email"] = nil
	session.Save(r, w)
	gothic.Logout(w, r)
}
