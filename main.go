package main

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type ErrorMessage struct {
	Code    string
	Message string
}

type JsonCredentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var cfg Config

// var jwtKey []byte
var jwtKey = []byte("my_secret_key")

func main() {

	cfg = *loadConfig("config.yaml")

	// jwtKey = []byte(cfg.Jwt.Key)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
