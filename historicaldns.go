package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

func historicaldns(work chan string, wg *sync.WaitGroup, lookupType string) {
	defer wg.Done()
	for text := range work {
		printAllPagesHistoricalDNS(text, lookupType)
	}
}

func printAllPagesHistoricalDNS(domain string, lookupType string) {
	response := getResponse(http.MethodGet, "history/"+domain+"/dns/"+lookupType, "")
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
		response = getResponse(http.MethodGet, "history/"+domain+"/dns/"+lookupType+"/?page="+strconv.Itoa(i), "")
		fmt.Println(response)
	}
}
