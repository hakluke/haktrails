package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
)

type Meta struct {
	Pages   int    `json:"total_pages"`
	Query   string `json:"query"`
	Page    int    `json:"page"`
	MaxPage int    `json:"max_page"`
}

func associatedDomains(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		printAllPages(text) // text is the domain
	}
}

func printAllPages(domain string) {
	response := getResponse("GET", "domain/"+domain+"/associated")
	fmt.Println(response)
	var results map[string]interface{}
	json.Unmarshal([]byte(response), &results)
	metaInterface, _ := results["meta"].(map[string]interface{})
	var maxPage float64
	if metaInterface["max_page"] == nil {
		if output == "list" {
			return
		} else {
			fmt.Println(response)
		}
		return
	}

	maxPage = metaInterface["max_page"].(float64) // total number of pages

	if output == "list" {
		parseAndPrintDomains(response) // print the first page
		// print all the other pages
		for i := 2; i <= int(maxPage); i++ {
			response = getResponse("GET", "domain/"+domain+"/associated?page="+strconv.Itoa(i))
			parseAndPrintDomains(response)
		}
	} else {
		fmt.Println(response) // print the first page
		// print all the other pages
		for i := 2; i <= int(maxPage); i++ {
			response = getResponse("GET", "domain/"+domain+"/associated?page="+strconv.Itoa(i))
			fmt.Println(response)
		}
	}
}

func parseAndPrintDomains(body string) {
	var results map[string]interface{}
	json.Unmarshal([]byte(body), &results)
	recordInterfaces := results["records"].([]interface{})
	for _, record := range recordInterfaces {
		recordMap := record.(map[string]interface{})
		fmt.Println(recordMap["hostname"].(string))
	}
}
