package accountapi

import (
	"encoding/json"
	"entities"
	"fmt"

	//"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "MySecretKey"

func GenerateToken(response http.ResponseWriter, request *http.Request) {
	var account entities.Account
	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": account.Username,
			"password": account.Password,
			"exp":      time.Now().Add(time.Minute * 5).Unix(), //thời gian sống là 5 phút
		})
		tokenString, err2 := token.SignedString([]byte(secretKey))
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, entities.JWTToken{Token: tokenString})
		}
	}
}

func CheckToken(response http.ResponseWriter, request *http.Request) {
	tokenString := request.Header.Get("key")
	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err == nil && result.Valid {
		fmt.Println("Xác thực thành công")
	} else {
		fmt.Println("xác thực không thành công")
	}
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	claims := jwt.MapClaims{} //sử dụng interface map để giải mã chuỗi string
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil 
	})
	if err != nil {
		return nil, false
	}
	fmt.Println(claims)
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}

	return claims, true
}

func GetInfoFromToken(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("key") //get token dựa trên request header
	fmt.Println(tokenStr)
	claim, _ := ExtractClaims(tokenStr) //lấy ra claim
	fmt.Println(claim)
	jsonString, _ := json.Marshal(claim)
	fmt.Println(jsonString)
	respondWithJSON(w, http.StatusOK, string(jsonString))
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
