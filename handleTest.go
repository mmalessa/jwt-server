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

func handleTest(w http.ResponseWriter, r *http.Request) {

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

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	claimsJson, _ := json.Marshal(claims)
	fmt.Fprint(w, string(claimsJson))
	log.Printf("TEST: (%s) Token expires at: %s\n", claims.Username, time.Unix(claims.ExpiresAt, 0))
}
