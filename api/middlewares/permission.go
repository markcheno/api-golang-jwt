package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	model "../models"
)

//Flags for biwise
const (
	FLAGREAD   int = 1  // 000001
	FLAGWRITE  int = 2  // 000010
	FLAGUPDATE int = 4  // 000100
	FLAGDELETE int = 8  // 001000
	FLAGALL    int = 16 // 010000
	GOD        int = 32 // 100000
)

// CheckFlags, true for correspond action
func checkFlags(route string, perms int) bool {
	var flag = false
	route = strings.ToUpper(route)

	switch route {
	case "GET":

		if (perms & FLAGREAD) > 0 {
			flag = true
		} else if (perms & FLAGALL) > 0 {
			flag = true
		}
		break
	case "POST":

		if (perms & FLAGWRITE) > 0 {
			flag = true
		} else if (perms & FLAGALL) > 0 {
			flag = true
		}

		break
	case "UPDATE":

		if (perms & FLAGUPDATE) > 0 {
			flag = true
		} else if (perms & FLAGALL) > 0 {
			flag = true
		}

		break
	case "DELETE":

		if (perms & FLAGDELETE) > 0 {
			flag = true
		} else if (perms & FLAGALL) > 0 {
			flag = true
		}

		break
	case "ALL":
		if (perms & FLAGALL) > 0 {
			flag = true
		}

		break
	case "GOD":
		if (perms & FLAGALL) > 0 {
			flag = true
		}

		break
	}

	return flag
}

// CheckUserPermission, get a method used and endpoint and check
// if user have permissions granted
func checkUserPermisson(action string, endpoint string, claims model.Claims) bool {
	//recover the user_id from context

	log.Printf("[CheckUserPermission] UserID = %s", claims.UserID)

	//find no mongodb { user: user_id, endpoint: endpoint }
	//MOC return mongodb the actions
	var actions = map[string]int{
		"GET":    1,
		"POST":   2,
		"UPDATE": 4,
		"DELETE": 8,
		//"ALL":    16,
		//"GOD":    32,
	}

	//intereate each permission and sum bit each bit
	var perms = 0
	for _, v := range actions {
		//fmt.Printf("method[%s] bit[%d]\n", k, v)
		perms |= v
		//fmt.Printf("permission = %d\n", perms)
	}

	//check if return something and the flags are ok
	if perms > 0 && checkFlags(action, perms) {
		log.Printf("[CheckUserPermission] User Authorized")
		return true
	}

	log.Printf("[CheckUserPermission] User Unauthorized")
	return false
}

// RequirePermission middleware for validate permission of user
func RequirePermission(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claims, ok := r.Context().Value(model.MyKey).(model.Claims)
		if !ok {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"message":"Error on decode Context JWT"}`)
			return
		}

		log.Printf("[RequirePermisson] method=%s EndPoint=%s", r.Method, r.URL.RequestURI())

		if checkUserPermisson(r.Method, r.URL.RequestURI(), claims) {
			next.ServeHTTP(w, r)
		}

		w.WriteHeader(http.StatusUnauthorized)
	})
}
