package main

import (
	"fmt"
	"sync"
)

func company(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		response := getResponse("GET", "company/"+text, "")
		fmt.Println(response)
	}
}
