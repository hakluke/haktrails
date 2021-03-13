package main

import (
	"fmt"
	"sync"
)

func associatedDomains(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		fmt.Println(text)
	}
}
