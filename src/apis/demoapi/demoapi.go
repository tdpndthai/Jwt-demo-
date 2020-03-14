package demoapi

import (
	"fmt"
	"net/http"
)

func Demo1(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "demo 1 api")
}

func Demo2(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "demo 2 api")
}
