package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func usage() {

	//color start
	color.Set(color.FgYellow)
	response := getResponse(http.MethodGet, "account/usage", "")
	fmt.Println(response)
	//color stop
	color.Unset()
}
