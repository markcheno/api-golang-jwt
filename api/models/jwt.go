package models

import jwt "github.com/dgrijalva/jwt-go"

//Claims jwt
type Claims struct {
	UserID string `json:"user_id"`
	Admin  bool   `json:"admin"`
	// recommended having
	jwt.StandardClaims
}

//TokenEndExpire return a token and expire timestamp
type TokenEndExpire struct {
	Token  string
	Expire string
}

// Key model for JWT authorization
type Key int

//MyKey jwt handdler
const (
	MyKey     Key = 0
	SecretKey     = "My-temporary-SECRETKey-2017"
)
