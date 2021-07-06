package main

import (
	"fmt"
	"sync"
)

// the main subdomains function will wait until it has received $buffer number of subdomains from stdin, then submit them
func submit(work chan string, wg *sync.WaitGroup, buffer int) {
	defer wg.Done()
	var subs []string
	for text := range work {
		subs = append(subs, text)
		if len(subs) >= buffer {
			submitSubdomains(subs)
			subs = nil
		}
	}
}

// get the subdomains + print them
func submitSubdomains(subs []string) {

	// convert the string slice into a list of strings
	var postBody string
	for c, s := range subs {
		if c == 0 {
			postBody = s
		} else {
			postBody = postBody + fmt.Sprintf("\n%s", s)
		}
	}

	// send it
	fmt.Println("Response:", getResponse("POST", "submit/hostnames", postBody))
}
