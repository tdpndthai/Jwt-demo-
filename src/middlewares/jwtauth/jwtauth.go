package jwtauth

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "MySecretKey"

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		tokenString := request.Header.Get("key")
		if tokenString == "" {
			respondWithError(response, http.StatusUnauthorized, "chưa có quyền")
		} else {
			result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})
			if err == nil && result.Valid {
				next.ServeHTTP(response, request)
			} else {
				respondWithError(response, http.StatusUnauthorized, "chưa có quyền")
			}
		}
	})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
