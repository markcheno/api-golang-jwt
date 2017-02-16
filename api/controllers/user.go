package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pressly/chi"
	"gopkg.in/mgo.v2/bson"

	db "../dbs"
	model "../models"
)

//GetUser get a user by Id
func GetUser(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ss := s.MongoDB.Copy()
		defer ss.Close()

		// Grab id
		id := chi.URLParam(r, "id")

		// Verify id is ObjectId, otherwise bail
		if !bson.IsObjectIdHex(id) {
			w.WriteHeader(404)
			return
		}
		// Grab id
		oid := bson.ObjectIdHex(id)
		// Stub user
		u := model.User{}
		// Fetch user
		if err := ss.DB("login").C("users").FindId(oid).One(&u); err != nil {
			w.WriteHeader(404)
			return
		}

		// Marshal provided interface into JSON structure
		uj, _ := json.Marshal(u)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", uj)
	}
}

//CreateUser create a new user
func CreateUser(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ss := s.MongoDB.Copy()
		defer ss.Close()

		// Stub an user to be populated from the body
		u := model.User{}

		// Populate the user data
		json.NewDecoder(r.Body).Decode(&u)

		// Add an Id
		u.ID = bson.NewObjectId()

		// Write the user to mongo
		ss.DB("login").C("users").Insert(u)

		// Marshal provided interface into JSON structure
		uj, _ := json.Marshal(u)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", uj)
	}
}
