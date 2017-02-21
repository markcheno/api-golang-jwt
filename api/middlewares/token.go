package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	mid "../models"
)

// TokenAuthentication middleware for vaidation token
func TokenAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var tokenStr string

		// Get token from query params
		tokenStr = r.URL.Query().Get("jwt")

		// Get token from authorization header
		if tokenStr == "" {
			bearer := r.Header.Get("Authorization")
			if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
				tokenStr = bearer[7:]
			}
		}

		// Get token from cookie
		if tokenStr == "" {
			cookie, err := r.Cookie("jwt")
			if err == nil {
				tokenStr = cookie.Value
			}
		}

		//if token not recovered on the last 3 steps
		if tokenStr == "" {
			log.Printf("[RequireTokenAuthentication] Not have Token to validate")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(tokenStr, &mid.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(mid.SecretKey), nil
		})

		if claims, ok := token.Claims.(*mid.Claims); ok && token.Valid {
			log.Printf("[RequireTokenAuthentication] Token valid! Go forward")
			ctx := context.WithValue(r.Context(), mid.JwtKey, *claims)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else if ve, ok := err.(*jwt.ValidationError); ok {

			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				log.Printf("[RequireTokenAuthentication] Thats not even token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				log.Printf("[RequireTokenAuthentication] Token is expired or not active")
			} else {
				log.Printf("[RequireTokenAuthentication] Couldn't handle token: %q", err)
			}

		} else {
			log.Printf("[RequireTokenAuthentication] Couldn't handle token: %q", err)
		}

		w.WriteHeader(http.StatusUnauthorized)
	})
}
