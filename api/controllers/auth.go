package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "../dbs"
	model "../models"
	service "../services"
)

//Home a home API
func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Public routes!"))
	}
}

//Auth get a valid token and expire
func Auth(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user model.User
		decoder := json.NewDecoder(r.Body)
		errDecoder := decoder.Decode(&user)

		if errDecoder != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, `{"message":"Incorrect Decode JSON on body"}`)
			return
		}

		t, err := service.GenerateToken(s, user)
		if err != nil {
			log.Printf("Error : %q", err)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, `{"message": %q}`, err)
			return
		}

		//write json
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"token":%q, "expire":%s}`, t.Token, t.Expire)
	}
}
