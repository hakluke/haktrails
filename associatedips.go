package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

func associatedIPs(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for text := range work {
		response := getResponse("GET", "company/"+text+"/associated-ips")
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
	recordInterfaces := results["records"].([]interface{})
	for _, record := range recordInterfaces {
		recordMap := record.(map[string]interface{})
		fmt.Println(recordMap["cidr"].(string))
	}
}
