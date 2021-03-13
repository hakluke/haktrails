package main

import (
	"fmt"
)

func usage() {
	response := getResponse("GET", "account/usage")
	fmt.Println(response)
}
