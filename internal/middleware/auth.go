package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const userIDKey ctxKey = "user_id"

// RequireAuth возвращает обёртку для защищённых хендлеров.
// secret — ваш []byte секрет для HS256.
func RequireAuth(secret []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if !strings.HasPrefix(auth, "Bearer ") {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			tokenStr := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer"))
			parsed, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				return secret, nil
			})
			if err != nil || !parsed.Valid {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			cs, ok := parsed.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			// ожидаем sub как число (id INTEGER), но поддержим и строку
			var uid int
			switch v := cs["sub"].(type) {
			case float64:
				uid = int(v)
			case string:
				if n, err := strconv.Atoi(v); err == nil {
					uid = n
				}
			}
			if uid == 0 {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), userIDKey, uid)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// UserID извлекает user_id, положенный мидлваркой.
func UserID(ctx context.Context) (int, bool) {
	v, ok := ctx.Value(userIDKey).(int)
	return v, ok && v > 0
}
