package main

import (
        "encoding/json"
        "io/ioutil"
        "log"
        "net/http"
        "strings"
)

// get a response from an endpoint as a string
func getResponse(method string, url string, postBody string) string {
        for {
                response, status := tryRequest(method, url, postBody)
                errorEncountered, message := checkJSONError(response, status)
                if errorEncountered {
                        return "JSON error: " + message
                } else {
                        return string(response)
                }
        }
}

func tryRequest(method string, url string, postBody string) (string, int) {
        client := &http.Client{}
        url = apiEndpoint + url
        req, err := http.NewRequest(method, url, strings.NewReader(postBody))
        if err != nil {
                log.Println("Error creating request:", err)
                return "", 0
        }
        req.Header.Set("apikey", config.SecurityTrails.Key)
        resp, err := client.Do(req)

        if err != nil {
                log.Println("Error retrieving:", url, err)
                return "", 0
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)

        if err != nil {
                log.Println("Error reading request body.", err)
                return "", 0
        }

        return string(body), resp.StatusCode
}

// Checks for {"message":"API rate limit exceeded"} or similar
func checkJSONError(body string, status int) (bool, string) {
        var results map[string]string
        err := json.Unmarshal([]byte(body), &results)
        if (err != nil && status == 200) {
                return false, ""
        }
        return true, results["message"]
}
