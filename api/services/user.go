package services

import (
	"gopkg.in/mgo.v2/bson"

	"fmt"

	db "../dbs"
	model "../models"
	lib "../shared"
)

var session = db.StartMongoDB("Middleware / User Service").Session

//UserIsValidOnProject seek the user on the project profile
func UserIsValidOnProject(slug string, userID string) (model.Project, error) {

	ss := session.Copy()
	defer ss.Close()

	oid := bson.ObjectIdHex(userID)

	//find user
	u := model.Project{}
	if err := ss.DB("login").C("projects").Find(bson.M{"slug": slug, "users": oid}).One(&u); err != nil {
		return u, fmt.Errorf("User not valid on project")
	}

	return u, nil
}

// UserGetPermissions return a permisson of user by project and endpoint
func UserGetPermissions(userID string, projectID string, endpoint string) (model.Permission, error) {

	ss := session.Copy()
	defer ss.Close()

	//change objectId
	endp := lib.GetRootEndpoint(endpoint)
	oid := bson.ObjectIdHex(userID)
	oidp := bson.ObjectIdHex(projectID)

	//find user
	u := model.Permission{}
	if err := ss.DB("login").C("permissions").Find(bson.M{"owner": oid, "project": oidp, "endpoint": endp}).One(&u); err != nil {
		return u, fmt.Errorf("Permission not found")
	}

	return u, nil
}
