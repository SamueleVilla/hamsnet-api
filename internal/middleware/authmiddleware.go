package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/samuelevilla/hasnet-api/internal/httputil"
	"github.com/samuelevilla/hasnet-api/internal/types"
)

// HTTP middleware setting a value on the request context
func Auth(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				// get the value of the "Authorization" header from the request

				log.Println("Auth middleware running")

				header := r.Header.Get("Authorization")
				if header == "" {
					httputil.WriteError(w, http.StatusUnauthorized, "missing token")
					return
				}
				tokenString := strings.Replace(header, "Bearer ", "", 1)

				// verify token
				claims := &types.JwtClaims{}
				token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}

					return []byte(secret), nil
				})
				if err != nil {
					httputil.WriteError(w, http.StatusUnauthorized, "invalid token")
					return
				}

				if !token.Valid {
					httputil.WriteError(w, http.StatusUnauthorized, "invalid token")
					return
				}

				log.Println("Auth middleware setting user_id context value to", claims.UserId)
				ctx := httputil.ContextWithUser(r.Context(), &types.User{
					Id:       claims.UserId,
					Username: claims.Username,
					Email:    claims.Email,
					Roles:    claims.Roles,
				})
				next.ServeHTTP(w, r.WithContext(ctx))
			},
		)
	}
}
