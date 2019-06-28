package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func Welcome(w http.ResponseWriter, r *http.Request) {

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

	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, fmt.Sprintf("Welcome %s!", claims.Username))
	// log.Printf("WELCOME: Token expires at: %s\n", time.Unix(claims.ExpiresAt, 0))
}
