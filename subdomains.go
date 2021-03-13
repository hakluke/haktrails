package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

// the main subdomains function
func subdomains(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		retrieveAndPrintSubdomains(text)
	}
}

// get the subdomains + print them
func retrieveAndPrintSubdomains(domain string) {
	response := getResponse("GET", "domain/"+domain+"/subdomains")
	if output == "json" {
		fmt.Println(response)
	} else {
		parseAndPrintSubdomains(response, domain)
	}
}

// list the subdomains
func parseAndPrintSubdomains(body string, domain string) {
	var results map[string]interface{}
	json.Unmarshal([]byte(body), &results)
	subdomainInterfaces := results["subdomains"].([]interface{})
	for _, subdomain := range subdomainInterfaces {
		fmt.Println(subdomain.(string) + "." + domain)
	}
}
