package src

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/mrspec7er/license-request/service/user/internal/db"
)

type AuthMiddleware struct {
	Service AuthService
}

func (m AuthMiddleware) Authorize(roles ...string) func(http.Handler) http.Handler {
	return (func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var userEmail string
			err := m.Service.GetUserEmail(r, &userEmail)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
				return
			}

			user := &db.User{Email: userEmail}

			status, err := m.Service.FindUser(user)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(status)
				json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
				return
			}

			if !slices.Contains(roles, user.Role) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(map[string]string{"message": "Unauthorize user"})
				return
			}

			ctx := context.WithValue(r.Context(), UserContextKey, user)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	})
}
