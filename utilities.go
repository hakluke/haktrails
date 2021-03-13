package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// get a response from an endpoint as a string
func getResponse(method string, url string) string {
	for {
		response := tryRequest(method, url)
		errorEncountered, message := checkJSONError(response)
		if errorEncountered {
			log.Println("JSON error in response:", message)
		} else {
			return string(response)
		}
	}
}

func tryRequest(method string, url string) string {
	client := &http.Client{}
	url = apiEndpoint + url
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Set("apikey", config.SecurityTrails.Key)
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error retrieving:", url, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading request body.", err)
	}

	return string(body)
}

// Checks for {"message":"API rate limit exceeded"} or similar
func checkJSONError(body string) (bool, string) {
	var results map[string]string
	err := json.Unmarshal([]byte(body), &results)
	if err != nil {
		return false, ""
	}
	return true, results["message"]
}
