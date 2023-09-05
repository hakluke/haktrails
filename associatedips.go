package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func associatedIPs(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		response := getResponse(http.MethodGet, "company/"+text+"/associated-ips", "")
		if output == "list" {
			parseAndPrintIPs(response)
		} else {
			fmt.Println(response)
		}
	}
}

func parseAndPrintIPs(body string) {
	var results map[string]interface{}
	json.Unmarshal([]byte(body), &results)
	val, ok := results["records"]
	// Check if there are records in the request to avoid the interface conversion panic
	if !ok || val == nil {
		fmt.Println(body)
		return
	}
	recordInterfaces := results["records"].([]interface{})
	for _, record := range recordInterfaces {
		recordMap := record.(map[string]interface{})
		fmt.Println(recordMap["cidr"].(string))
	}
}
