package main

import (
	"fmt"
	"net/http"
	"sync"
)

func whois(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		response := getResponse(http.MethodGet, "domain/"+text+"/whois", "")
		fmt.Println(response)
	}
}
