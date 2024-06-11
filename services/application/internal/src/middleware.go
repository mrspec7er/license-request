package src

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/mrspec7er/license-request/services/application/internal/db"
)

type Middleware struct {
	Cache db.CacheRepository[*dto.User]
}

func (m Middleware) Authorize(roles ...string) func(http.Handler) http.Handler {
	return (func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authKey, err := m.GetUserAuthKey(r)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
				return
			}

			user := dto.User{}
			err = m.RetrieveUserSessions(w, r, authKey, &user)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
				return
			}

			if len(roles) > 0 && !slices.Contains(roles, user.Role) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(map[string]string{"message": "Unauthorize user"})
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
	err := m.Cache.Retrieve(context.Background(), key, user)
	if err != nil {
		return err
	}

	return nil
}
