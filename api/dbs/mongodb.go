package dbs

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
)

func start() {
	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{""},
		Timeout:  60 * time.Second,
		Database: "",
		Username: "",
		Password: "",
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	log.Printf("%s", mongoSession.LiveServers())
}
