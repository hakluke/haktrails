package main

import (
	"fmt"

	"github.com/fatih/color"
)

func usage() {

	//color start
	color.Set(color.FgYellow)
	response := getResponse("GET", "account/usage", "")
	fmt.Println(response)
	//color stop
	color.Unset()
}
