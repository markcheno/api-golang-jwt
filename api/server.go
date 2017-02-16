package main

import (
	"fmt"
	"net/http"

	route "./routes"
)

func main() {
	addr := ":3333"
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, route.Router())
}
