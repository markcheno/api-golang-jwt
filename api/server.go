package main

import (
	"fmt"
	"net/http"

	"log"

	db "./dbs"
	route "./routes"
	lib "./webapp"
)

func main() {
	sessions := db.StartDispatch()

	log.Printf(lib.GetPath())

	addr := ":3333"
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, route.Router(sessions))
}
