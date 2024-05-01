package src

import (
	"context"
	"errors"
	"net/http"
	"slices"

	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/mrspec7er/license-request-utility/response"
)

type Middleware struct {
	Util     *Util
	Response response.ResponseJSON
}

func (m Middleware) Authorize(roles ...string) func(http.Handler) http.Handler {
	return (func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authKey, err := m.GetUserAuthKey(r)

			if err != nil {
				m.Response.UnauthorizeUser(w)
				return
			}

			user := dto.User{}
			err = m.RetrieveUserSessions(w, r, authKey, &user)

			if err != nil {
				m.Response.GeneralErrorHandler(w, 403, errors.New("expired token"))
				return
			}

			if len(roles) > 0 && !slices.Contains(roles, user.Role) {
				m.Response.GeneralErrorHandler(w, 403, errors.New("user access denied"))
				return
			}

			ctx := context.WithValue(r.Context(), dto.UserContextKey, user)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	})
}

func (m Middleware) GetUserAuthKey(r *http.Request) (string, error) {
	cookie, err := r.Cookie(string(dto.AuthCookieName))
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func (m Middleware) RetrieveUserSessions(w http.ResponseWriter, r *http.Request, key string, user *dto.User) error {
	err := m.Util.MemcacheRetrieve(context.Background(), key, &user)
	if err != nil {
		return err
	}

	return nil
}
