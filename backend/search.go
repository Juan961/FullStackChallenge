package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Search(credentials string, term string) (results []byte) {
	// Define search parameters
	params := map[string]interface{}{
		"search_type": "match",
		"query": map[string]string{
			"term": term,
		},
		"fields": []string{"_all"},
	}

	// Define request headers
	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Basic " + credentials,
	}

	// Define the index and host URL
	index := "Emails"
	zincHost := "http://localhost:4080"
	zincURL := zincHost + "/api/" + index + "/_search"

	// Marshal the parameters to JSON
	jsonData, err := json.Marshal(params)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", zincURL, strings.NewReader(string(jsonData)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the request headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request using an HTTP client
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	// Decode the response body into a map
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	// Initialize the response map
	response := make(map[string]interface{})

	// Accessing "hits" -> "hits" array
	if hits, ok := result["hits"].(map[string]interface{}); ok {
			if hitsArray, ok := hits["hits"].([]interface{}); ok {
					response["hits"] = hitsArray
			} else {
					fmt.Println("Error: 'hits' is not a slice")
			}
	} else {
			fmt.Println("Error: 'hits' is not a map")
	}

	// Accessing "took"
	if took, ok := result["took"].(float64); ok {
			response["took"] = took
	} else {
			fmt.Println("Error: 'took' is not a float64")
	}

	// Accessing "total" (total value)
	if hits, ok := result["hits"].(map[string]interface{}); ok {
	    if total, ok := hits["total"].(map[string]interface{}); ok {
	        if value, ok := total["value"].(float64); ok {
	            response["total"] = value
	        } else {
	            fmt.Println("Error: 'value' is not a float64")
	        }
	    } else {
	        fmt.Println("Error: 'total' is not a map")
	    }
	}

	results, err = json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	return results
}
