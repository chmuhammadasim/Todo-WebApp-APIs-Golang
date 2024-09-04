// middleware/auth.go
package middleware

import (
	"context"
	"net/http"
	"strings"
	"todo-app/utils"

	"github.com/dgrijalva/jwt-go"
)

// Role-based middleware to check for required role
func AuthMiddleware(requiredRole string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, nil
				}
				return []byte("your_secret_key"), nil
			})
			if err != nil || !token.Valid {
				utils.SendError(w, http.StatusUnauthorized, "Invalid token")
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				utils.SendError(w, http.StatusUnauthorized, "Invalid token")
				return
			}

			// Check role
			role, ok := claims["role"].(string)
			if !ok || role != requiredRole {
				utils.SendError(w, http.StatusForbidden, "Insufficient role")
				return
			}

			ctx := context.WithValue(r.Context(), "user", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
