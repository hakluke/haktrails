package main

import (
	"fmt"
	"sync"
)

func details(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		response := getResponse("GET", "domain/"+text)
		fmt.Println(response)
	}
}
