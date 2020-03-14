package main

import (
	"apis/accountapi"
	"apis/demoapi"
	"fmt"
	"middlewares/jwtauth"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	//Creating the HTTP server use mux
	router.HandleFunc("/api/account/createacc", accountapi.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/account/updateacc", accountapi.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/account/generatekey", accountapi.GenerateToken).Methods(http.MethodPost)
	router.HandleFunc("/api/account/checktoken", accountapi.CheckToken).Methods(http.MethodGet)
	router.Handle("/api/demo/demo1", jwtauth.JWTAuth(http.HandlerFunc(demoapi.Demo1))).Methods(http.MethodGet)
	router.Handle("/api/demo/demo2", jwtauth.JWTAuth(http.HandlerFunc(demoapi.Demo2))).Methods(http.MethodGet)
	router.HandleFunc("/api/account/getinfo", accountapi.GetInfoFromToken).Methods(http.MethodGet)

	err := http.ListenAndServe(":4000", router)
	if err != nil {
		fmt.Println(err)
	}
}
