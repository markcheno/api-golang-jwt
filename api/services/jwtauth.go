package services

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	model "../models"
)

//GenerateToken for a valid user
func GenerateToken(user model.User) (model.TokenEndExpire, error) {
	// moc to return
	var te model.TokenEndExpire
	// expire unix
	var expir = time.Now().Add(time.Minute * 60).Unix()

	// Connect mongodb and check user and password
	if user.Email != "tzilli@inviron.com.br" && user.Password != "123" {
		return te, fmt.Errorf("Error : Invalid User or Password")
	}

	//claims HARDCODE
	//change this
	claims := model.Claims{
		"1234567890",
		true,
		jwt.StandardClaims{
			ExpiresAt: expir,
			Issuer:    "localhost:3333",
		},
	}

	// create token & sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(model.SecretKey))
	if err != nil {
		return te, fmt.Errorf("Error :%q", err)
	}

	te.Token = t
	te.Expire = fmt.Sprintf("%q", time.Unix(expir, 0))

	return te, nil
}
