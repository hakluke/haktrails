package main

import (
	"fmt"
	"sync"
)

func tags(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		response := getResponse("GET", "domain/"+text+"/tags")
		fmt.Println(response)
	}
}
