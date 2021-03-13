package main

import (
	"fmt"
	"sync"
)

func associatedips(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		response := getResponse("GET", "company/"+text+"/associated-ips")
		fmt.Println(response)
	}
}
