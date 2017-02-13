package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/rs/cors"
)

//Claims jwt
type Claims struct {
	UserID string `json:"user_id"`
	Admin  bool   `json:"admin"`
	// recommended having
	jwt.StandardClaims
}

// User model
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Key model for JWT authorization
type Key int

// Constants
const (
	ValidEmail     = "tzilli@inviron.com.br"
	ValidPass      = "123"
	SecretKey      = "My,Secrety,Key,Of,Inviron"
	MyKey      Key = 0
)

func main() {
	addr := ":3333"
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, router())
}

// RequireTokenAuthentication middleware for vaidation token
func RequireTokenAuthentication(next http.Handler) http.Handler {
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

		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			log.Printf("Token valid! Go forward")
			ctx := context.WithValue(r.Context(), MyKey, *claims)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else if ve, ok := err.(*jwt.ValidationError); ok {

			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				log.Printf("Thats not even token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				log.Printf("Token is expired or not active")
			} else {
				log.Printf("Couldn't handle token: %q", err)
			}

		} else {
			log.Printf("Couldn't handle token: %q", err)
		}

		w.WriteHeader(http.StatusUnauthorized)
	})
}

func router() http.Handler {
	r := chi.NewRouter()

	//CORS setup
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(cors.Handler)
		r.Use(RequireTokenAuthentication)

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(MyKey).(Claims)
			if !ok {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(500)
				fmt.Fprintf(w, `{"message":"Error on decode Context JWT"}`)
				return
			}
			w.Write([]byte(fmt.Sprintf("protected area. USER ID = %v", claims.UserID)))
		})
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(cors.Handler)

		// home
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Public routes!"))
		})

		// Authenticate user
		r.Post("/auth", func(w http.ResponseWriter, r *http.Request) {

			var user User
			decoder := json.NewDecoder(r.Body)
			errDec := decoder.Decode(&user)

			if errDec != nil {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(201)
				fmt.Fprintf(w, `{"message":"Incorrect Decode JSON on body"}`)
				return
			}

			// Connect mongodb and check user and password
			//if the email and pass valid.
			if user.Email != ValidEmail && user.Password != ValidPass {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(201)
				fmt.Fprintf(w, `{"message":"User invalid"}`)
				return
			}

			claims := Claims{
				"super-id-of-mongodb-user",
				true,
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Second * 3600 * 24).Unix(),
					Issuer:    "localhost:3333",
				},
			}
			// create token & sign token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			t, err := token.SignedString([]byte(SecretKey))
			if err != nil {
				log.Printf("Error generate token : %q", err)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(201)
				fmt.Fprintf(w, `{"message": "Error on generate token"}`)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(201)
			fmt.Fprintf(w, `{"token":%q}`, t)
		})

	})

	return r
}
