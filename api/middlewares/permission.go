package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	model "../models"
	service "../services"
)

//Flags for biwise
const (
	FLAGREAD   int = 1  // 000001
	FLAGWRITE  int = 2  // 000010
	FLAGUPDATE int = 4  // 000100
	FLAGDELETE int = 8  // 001000
	OWNER      int = 16 // 010000
	ADMIN      int = 32 // 100000
)

// CheckFlags, true for correspond action
func checkFlags(route string, perms int) bool {
	var flag = false
	route = strings.ToUpper(route)

	switch route {
	case "GET":

		if (perms & FLAGREAD) > 0 {
			flag = true
		} else if (perms & OWNER) > 0 {
			flag = true
		}
		break
	case "POST":

		if (perms & FLAGWRITE) > 0 {
			flag = true
		} else if (perms & OWNER) > 0 {
			flag = true
		}

		break
	case "UPDATE":

		if (perms & FLAGUPDATE) > 0 {
			flag = true
		} else if (perms & OWNER) > 0 {
			flag = true
		}

		break
	case "DELETE":

		if (perms & FLAGDELETE) > 0 {
			flag = true
		} else if (perms & OWNER) > 0 {
			flag = true
		}

		break
	case "OWNER":
		if (perms & OWNER) > 0 {
			flag = true
		}

		break
	case "ADMIN":
		if (perms & ADMIN) > 0 {
			flag = true
		}

		break
	}

	return flag
}

// CheckUserPermission, get a method used and endpoint and check
// if user have permissions granted
func checkUserPermisson(action string, endpoint string, projectID string, claims model.Claims) bool {

	log.Printf("[UserHavePermission] UserID = %q", claims.UserID)
	log.Printf("[UserHavePermission] ProjectID = %q", projectID)

	perm, perr := service.UserGetPermissions(claims.UserID, projectID, endpoint)
	if perr != nil {
		log.Printf("[UserHavePermission] Err: %s", perr)
		return false
	}

	log.Printf("[UserHavePermission] = %+v \n", perm)

	//interate each permission and sum, bit each bit
	var perms = 0
	for _, vv := range perm.Permission {
		perms |= vv.Value
	}

	//check if return something and the flags are ok
	if perms > 0 && checkFlags(action, perms) {
		log.Printf("[UserHavePermission] User Authorized")
		return true
	}

	log.Printf("[UserHavePermission] User Unauthorized")
	return false
}

// UserHavePermission middleware for validate permission of user
func UserHavePermission(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claims, ok := r.Context().Value(model.JwtKey).(model.Claims)
		if !ok {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"message":"Error on decode Context JWT"}`)
			return
		}

		projectID, ok := r.Context().Value(model.ProjKey).(string)
		if !ok {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"message":"Error on decode Context ProjectID"}`)
			return
		}

		// Check in future
		// mgos, ok := r.Context().Value(model.DbKey).(*mgo.Database)
		// if !ok {
		// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 	w.WriteHeader(500)
		// 	fmt.Fprintf(w, `{"message":"Error on decode Session MongoDB"}`)
		// 	return
		// }

		log.Printf("[UserHavePermission] method=%s EndPoint=%s", r.Method, r.URL.RequestURI())

		if checkUserPermisson(r.Method, r.URL.RequestURI(), projectID, claims) {
			next.ServeHTTP(w, r)
		}

		w.WriteHeader(http.StatusUnauthorized)
		return
	})
}
