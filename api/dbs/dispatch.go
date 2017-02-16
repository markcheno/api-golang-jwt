package dbs

import (
	"gopkg.in/mgo.v2"
)

//Dispatch choose a db session
// add new dispath of other database just put here
// the session.
type Dispatch struct {
	MongoDB *mgo.Session
}

//StartDispatch load up connections
func StartDispatch() *Dispatch {
	//add session of mongodb
	mongosession := StartMongoDB().Session

	return &Dispatch{MongoDB: mongosession}

}
