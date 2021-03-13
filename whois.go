package main

import (
	"fmt"
	"sync"
)

func whois(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		response := getResponse("GET", "domain/"+text+"/whois")
		fmt.Println(response)
	}
}
