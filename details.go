package main

import (
	"fmt"
	"net/http"
	"sync"
)

func details(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		response := getResponse(http.MethodGet, "domain/"+text, "")
		fmt.Println(response)
	}
}
