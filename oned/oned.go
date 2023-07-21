package oned

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Oned() {
	url := "https://example.com/data.json" // Replace this with the actual URL of the JSON data

	// Make an HTTP GET request to fetch the JSON data
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer response.Body.Close()

	// Check if the response status code is 200 (OK)
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Server returned non-OK status: %d\n", response.StatusCode)
		return
	}

	// Decode the JSON data
	var data []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Now you can work with the 'data' slice of maps
	for index := range data {
		fmt.Printf("index: %d\n", index)
	}
}
