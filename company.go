package main

import (
	"fmt"
	"net/http"
	"sync"
)

func company(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		response := getResponse(http.MethodGet, "company/"+text, "")
		fmt.Println(response)
	}
}
