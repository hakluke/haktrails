package main

import (
	"fmt"
	"net/http"
	"sync"
)

func tags(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		response := getResponse(http.MethodGet, "domain/"+text+"/tags", "")
		fmt.Println(response)
	}
}
