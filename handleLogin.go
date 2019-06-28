package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials JsonCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		// Bad Request
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorMessage{Code: "400", Message: http.StatusText(400)})
		return
	}

	expectedPassword, ok := cfg.Credentials[credentials.Username]

	if !ok || expectedPassword != credentials.Password {
		// Unauthorized
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(ErrorMessage{Code: "401", Message: http.StatusText(401)})
		return
	}

	expiresAtTime := time.Now().Add(time.Duration(cfg.Jwt.ExpirationTime) * time.Minute)
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAtTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		// Internal Server Error
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ErrorMessage{Code: "500", Message: http.StatusText(500)})
		return
	}
	fmt.Fprint(w, tokenString)
	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// log.Printf("LOGIN:   Token expires at: %s\n", time.Unix(claims.ExpiresAt, 0))
}
