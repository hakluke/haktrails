package main

import (
	"fmt"
)

func ping() {
	response := getResponse("GET", "/ping")
	fmt.Println(response)
}
