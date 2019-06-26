package main

import (
	"log"
	"net/http"
)

var cfg Config

func main() {

	cfg = *loadConfig("config.yaml")

	http.HandleFunc("/login", Login)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
