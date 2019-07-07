package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func handleRefresh(w http.ResponseWriter, r *http.Request) {

	auth := r.Header.Get("Authorization")
	if auth == "" {
		// Bad Request
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorMessage{Code: "400", Message: http.StatusText(400)})
		return
	}

	tokenString := strings.TrimPrefix(auth, "Bearer ")

	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Println("INCORRECT TOKEN STRING")
		// Bad Request
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(ErrorMessage{Code: "401", Message: http.StatusText(401)})
		return
	}

	if !tkn.Valid {
		// Unauthorized
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(ErrorMessage{Code: "401", Message: http.StatusText(401)})
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			// Unauthorized
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorMessage{Code: "400", Message: http.StatusText(400)})
			return
		}
		// Bad Request
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(ErrorMessage{Code: "401", Message: http.StatusText(401)})
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expiresAtTime := time.Now().Add(time.Duration(cfg.Jwt.ExpirationTime) * time.Minute)

	claims.ExpiresAt = expiresAtTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// Internal Server Error
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ErrorMessage{Code: "500", Message: http.StatusText(500)})
		return
	}

	response := map[string]string{
		"token": newTokenString,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseM, _ := json.Marshal(response)
	fmt.Fprint(w, string(responseM))
	log.Printf("REFRESH: (%s) Token expires at: %s\n", claims.Username, time.Unix(claims.ExpiresAt, 0))
}
