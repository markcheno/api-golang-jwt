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
	JwtKey    Key = 100000
	DbKey     Key = 200000
	UserKey   Key = 300000
	ProjKey   Key = 400000
	SecretKey     = "My-temporary-SECRETKey-2017"
)
