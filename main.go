package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

var cfg *Config
var err error
var jwtKey []byte

func main() {

	log.Println("Starting JWT Test Server")

	configFile := "config.yaml"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	cfg, err = loadConfig(configFile)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: %v", err))
		return
	}

	log.Println(fmt.Sprintf("Server port:%d", cfg.Server.Port))

	jwtKey = []byte(cfg.Jwt.Key)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), nil))
}
