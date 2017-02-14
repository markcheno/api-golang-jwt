package main

import (
	"fmt"
	"log"
	"strings"
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
	}

	return flag
}

func checkUserPermisson(action string, endpoint string) {
	//recover the user_id from context

	// find no mongodb { user: user_id, endpoint: endpoint }

	//return mongodb the actions
	var actions = map[string]int{
		"GET":    1,
		"POST":   2,
		"UPDATE": 4,
		"DELETE": 8,
		//"ALL": 16,
	}

	//intereate each permission and sum
	var perms = 0
	for k, v := range actions {
		fmt.Printf("method[%s] bit[%d]\n", k, v)
		perms |= v
		fmt.Printf("permission = %d\n", perms)
	}

	//check if return something and the flags are ok
	if perms > 0 && checkFlags(action, perms) {
		log.Printf("Accept - Authorized")
	} else {
		log.Printf("Denied - UNAuthorized")
	}

}

func main() {
	checkUserPermisson("DELETE", "report")
}
