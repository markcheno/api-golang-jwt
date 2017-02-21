package dbs

import (
	"github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

//Dispatch choose a db session
// add new dispath of other database just put here
// the session.
type Dispatch struct {
	MongoDB *mgo.Session
	Logger  *logrus.Logger
}

//StartDispatch load up connections
func StartDispatch() *Dispatch {
	//add session of mongodb
	mongosession := StartMongoDB("Dispatch Service").Session
	// add logger for dispatch
	logger := Logger()

	return &Dispatch{MongoDB: mongosession, Logger: logger}

}
