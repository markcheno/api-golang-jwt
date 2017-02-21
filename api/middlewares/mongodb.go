package middlewares

import (
	"context"
	"log"
	"net/http"
	"time"

	model "../models"
	mgo "gopkg.in/mgo.v2"
)

// MongoMiddleware adds mgo MongoDB to context
func MongoMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("MongoDB on request!")

		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:   []string{"localhost:27017"},
			Timeout: 60 * time.Second,
			//Database: "",
			//Username: "",
			//Password: "",
		}

		mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
		if err != nil {
			log.Fatalf("[MongoDB] CreateSession: %s\n", err)
		}
		mongoSession.SetMode(mgo.Monotonic, true)

		rs := mongoSession.Clone()
		defer rs.Close()

		db := rs.DB("login")
		ctx := context.WithValue(r.Context(), model.DbKey, db)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
