package dbs

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
)

//MgoSession and session
type MgoSession struct {
	Session *mgo.Session
}

func newMgoSession(s *mgo.Session) *MgoSession {
	return &MgoSession{s}
}

//StartMongoDB initialize session on mongodb
func StartMongoDB() *MgoSession {

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:   []string{"localhost:27017"},
		Timeout: 60 * time.Second,
		//Database: "",
		//Username: "",
		//Password: "",
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	mongoSession.SetMode(mgo.Monotonic, true)

	return newMgoSession(mongoSession)
}
