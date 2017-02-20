package services

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	jwt "github.com/dgrijalva/jwt-go"

	db "../dbs"
	model "../models"
	lib "../shared"
)

//GenerateToken for a valid user
func GenerateToken(s *db.Dispatch, user model.User) (model.TokenEndExpire, error) {

	ss := s.MongoDB.Copy()
	defer ss.Close()

	// moc to return
	var te model.TokenEndExpire
	// expire unix
	var expir = time.Now().Add(time.Minute * 60).Unix()

	//find user
	u := model.User{}
	if err := ss.DB("login").C("users").Find(bson.M{"email": user.Email}).One(&u); err != nil {
		return te, fmt.Errorf("Password not match or user not found")
	}

	if errPwd := lib.Compare(u.Password, user.Password); errPwd != nil {
		return te, fmt.Errorf("Invalid Password")
	}

	//build claims
	claims := model.Claims{
		u.ID.Hex(),
		u.Admin,
		jwt.StandardClaims{
			ExpiresAt: expir,
			Issuer:    "localhost:3333",
		},
	}

	// create token & sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(model.SecretKey))
	if err != nil {
		return te, fmt.Errorf("%q", err)
	}

	te.Token = t
	te.Expire = fmt.Sprintf("%q", time.Unix(expir, 0))

	return te, nil
}
