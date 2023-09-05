package main

import (
	"fmt"
	"net/http"
)

func ping() {
	response := getResponse(http.MethodGet, "/ping", "")
	fmt.Println(response)
}
