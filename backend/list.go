package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func List() (response []byte) {
	// Get user credentials from environment variables
	user := "admin" // os.Getenv("SEARCH_USER")
	password := "Complexpass#123" // os.Getenv("SEARCH_PASSWORD")
	if user == "" || password == "" {
		fmt.Println("Environment variables SEARCH_USER and SEARCH_PASSWORD must be set")
		return
	}
	creds := user + ":" + password
	// Encode credentials to base64
	bas64encodedCreds := base64.StdEncoding.EncodeToString([]byte(creds))

	// List index parameters, # of records to return
	params := map[string]interface{}{
		"size": 100,
	}

	// Define request headers
	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Basic " + bas64encodedCreds,
	}

	// Define the index and host URL
	zincHost := "http://localhost:4080"
	zincURL := zincHost + "/api/index"

	// Marshal the parameters to JSON
	jsonData, err := json.Marshal(params)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", zincURL, strings.NewReader(string(jsonData)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the request headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	// Assert that result["list"] is a slice of interfaces
	list, ok := result["list"].([]interface{})
	if !ok {
		fmt.Println("Error: result['list'] is not a slice")
		return
	}

	// Iterate over the list
	for _, item := range list {
		// Assert that item is a map[string]interface{}
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			fmt.Println("Error: item is not a map")
			continue
		}

		// Check if the "name" key exists and is equal to "Emails"
		if name, exists := itemMap["name"].(string); exists && name == "Emails" {
			results, err := json.Marshal(itemMap["stats"])
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)
				return
			}

			return results
		}
	}

	// Return error if index not found
	fmt.Println("Index not found")

	return nil
}
