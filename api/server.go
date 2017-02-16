package main

import (
	"fmt"
	"net/http"

	db "./dbs"
	route "./routes"
)

func main() {
	sessions := db.StartDispatch()

	addr := ":3333"
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, route.Router(sessions))
}
