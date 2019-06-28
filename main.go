package main

import (
	"fmt"
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

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var cfg Config
var jwtKey []byte

func main() {

	log.Println("Starting JWT Test Server")
	cfg = *loadConfig("config.yaml")
	log.Println(fmt.Sprintf("Server port:%d\n", cfg.Server.Port))

	jwtKey = []byte(cfg.Jwt.Key)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), nil))
}
