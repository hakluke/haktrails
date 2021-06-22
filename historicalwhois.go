package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
)

func historicalwhois(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		printAllPagesHistoricalWhois(text)
	}
}

func printAllPagesHistoricalWhois(domain string) {
	response := getResponse("GET", "history/"+domain+"/whois", "")
	var results map[string]interface{}
	json.Unmarshal([]byte(response), &results)
	maxPage, ok := results["pages"].(float64) // total number of pages
	if !ok {                                  // response is rong
		fmt.Println(response)
		return
	}

	fmt.Println(response) // print the first page
	// now print all the other pages
	for i := 2; i <= int(maxPage); i++ {
		response = getResponse("GET", "history/"+domain+"/whois/?page="+strconv.Itoa(i), "")
		fmt.Println(response)
	}
}
