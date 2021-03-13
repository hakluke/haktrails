package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
// curl --request GET --url https://api.securitytrails.com/v1/domain/trello.com/subdomains  --header 'apikey: {{key}}'
func retrieveAndPrintSubdomains(domain string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiEndpoint+"/domain/"+domain+"/subdomains", nil)
	req.Header.Set("apikey", config.SecurityTrails.Key)
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error retrieving subdomains endpoint:", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error retreiving request body.", err)
	}
	bodyString := string(body)
	//TODO handle pagination
	parseAndPrintSubdomains(bodyString, domain)
}

func parseAndPrintSubdomains(body string, domain string) {
	var results map[string]interface{}
	json.Unmarshal([]byte(body), &results)
	subdomainInterfaces := results["subdomains"].([]interface{})
	for _, subdomain := range subdomainInterfaces {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(subdomain.(string) + "." + domain)
	}
}
