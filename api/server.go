package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"log"

	db "./dbs"
	route "./routes"
	lib "./shared"
)

func init() {
	// load config file
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file! Try get a path...")
		if err2 := godotenv.Load(lib.GetPath() + "/.env"); err2 != nil {
			log.Printf("Fail...")
			os.Exit(1)
		}
	}
}

func main() {
	sessions := db.StartDispatch()
	addr := os.Getenv("API_URL")

	log.Printf("[Server] Path: %s", lib.GetPath())
	fmt.Printf("[Server] Starting server on %v\n", addr)
	http.ListenAndServe(addr, route.Router(sessions))
}
