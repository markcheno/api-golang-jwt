package main

import (
	"fmt"
	"log"
	"strings"
)

//Flags for biwise
const (
	FLAGREAD   int = 1  // 0001
	FLAGWRITE  int = 2  // 0010
	FLAGUPDATE int = 4  // 0100
	FLAGDELETE int = 8  // 1000
	FLAGALL    int = 16 // 10000
	GOD        int = 32 // 100000
)

func checkFlags(route string, perms int) bool {
	var flag = false

	switch strings.ToUpper(route) {
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
	//recupera do contexto o user_id
	// find no mongodb { user: user_id, endpoint: endpoint }
	//return do mongo as actions
	var actions = map[string]int{
		//"GET":  1,
		//"POST": 2,
		//"UPDATE": 4,
		"DELETE": 8,
		//"ALL": 16,
	}

	var perms = 0
	for k, v := range actions {
		//if strings.Contains(action, k) {
		fmt.Printf("key[%s] value[%d]\n", k, v)
		perms |= v
		fmt.Printf("permission  = %d\n", perms)
		//}
	}

	if perms > 0 && checkFlags(action, perms) {
		log.Printf("Authorized")
	} else {
		log.Printf("UNauthorized")
	}

}

func main() {
	checkUserPermisson("GET", "report")
}
